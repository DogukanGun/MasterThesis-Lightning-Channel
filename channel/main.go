package channel

import (
	"MasterThesis/bitcoin"
	"MasterThesis/logger"
	"MasterThesis/recorder"
	"MasterThesis/structs"
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/go-pdf/fpdf"
	"github.com/lightningnetwork/lnd/lnrpc"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func OpenChannel(destinationPubKey string, lncli lnrpc.LightningClient) {
	bobPubkey := destinationPubKey //"03fea3149e0afff6b948299b247eb3995c54105fa69da6a9d9dd425beb43df3342"
	openChannelRequest := lnrpc.OpenChannelRequest{
		NodePubkeyString:   bobPubkey,
		LocalFundingAmount: 10_000_000,
	}
	sync, err := lncli.OpenChannelSync(context.TODO(), &openChannelRequest)
	if err != nil {
		logger.LogE(err.Error())
	}
	logger.LogI(sync.String())
}

func CloseChannel(Txid string, lncli lnrpc.LightningClient) {
	channelPoint := lnrpc.ChannelPoint{
		FundingTxid: &lnrpc.ChannelPoint_FundingTxidStr{
			FundingTxidStr: Txid,
		},
		OutputIndex: 0,
	}
	closeChannelRequest := lnrpc.CloseChannelRequest{
		ChannelPoint: &channelPoint,
	}
	res, err := lncli.CloseChannel(context.TODO(), &closeChannelRequest)
	if err != nil {
		logger.LogE("Channel close: ", err)
	} else {
		logger.LogS("Channel close res: ", res)
	}
	closingTxHex := ""
	for {
		update, err := res.Recv()
		if err != nil {
			logger.LogE("Failed to receive close channel update: ", err)
		}
		if update.GetChanClose().Success {
			closingTxHex = hex.EncodeToString(update.GetChanClose().ClosingTxid)
			break
		}
	}
	metadata := "Closing transaction: " + closingTxHex + "\n"
	WriteToDatabase()
	hash := SaveToIpfs("messages.pdf")
	metadata += "The cid: " + hash
	// Configure Bitcoin RPC client
	rpcClient, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         os.Getenv("BitcoinHost"),
		User:         os.Getenv("BitcoinUser"),
		Pass:         os.Getenv("BitcoinPass"),
		HTTPPostMode: true,
		DisableTLS:   true,
	}, nil)
	if err != nil {
		logger.LogE("Failed to create Bitcoin RPC client: ", err)
	}
	defer rpcClient.Shutdown()
	if err := bitcoin.PublishMetadata(os.Getenv("PrevTransaction"), 0, metadata, rpcClient); err != nil {
		logger.LogE("Error modifying and broadcasting transaction: ", err)
	}

}

func SaveToIpfs(fileName string) string {
	url := "https://node.lighthouse.storage/api/v0/add"
	apiKey := os.Getenv("LIGHTHOUSE_KEY")

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		logger.LogE("Error opening file: %v", err)
	}
	defer file.Close()

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file field
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		logger.LogE("Error creating form file: %v", err)
	}

	// Copy the file content to the part
	_, err = io.Copy(part, file)
	if err != nil {
		logger.LogE("Error copying file content: %v", err)
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		logger.LogE("Error closing writer: %v", err)
	}

	// Create the HTTP request
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		logger.LogE("Error creating request: %v", err)
	}

	// Set headers
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer "+apiKey)

	// Make the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logger.LogE("Error making request: %v", err)
	}
	defer response.Body.Close()

	// Check response status
	if response.StatusCode != http.StatusOK {
		logger.LogE("Request failed. Status: %s", response.Status)
	}

	// Print response body
	fmt.Println("Response:", response.Status)
	// Read response body
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.LogE("Error reading response body: %v", err)
	}

	// Parse JSON response
	var uploadResp structs.UploadResponse
	err = json.Unmarshal(respBody, &uploadResp)
	if err != nil {
		logger.LogE("Error decoding JSON: %v", err)
	}

	// Print specific fields
	fmt.Printf("Name: %s\nHash: %s\nSize: %s\n", uploadResp.Name, uploadResp.Hash, uploadResp.Size)
	return uploadResp.Hash
}

func WriteToDatabase() {
	db := recorder.Connect()
	var messages []structs.Messages
	recorder.Get(db, "Messages", "", &messages)
	messagesAsStr := ""
	for _, message := range messages {
		messagesAsStr += "From" + message.Peer + "- Message: " + message.Message + "\n"
	}
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 100, messagesAsStr)
	err := pdf.OutputFileAndClose("messages.pdf")
	if err != nil {
		logger.LogE(err)
	}
}
