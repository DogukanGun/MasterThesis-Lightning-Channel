package channel

import (
	"MasterThesis/bitcoin"
	"MasterThesis/logger"
	"context"
	"encoding/hex"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/lightningnetwork/lnd/lnrpc"
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
	metadata := ""
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
	//Upload texts to lighthouse =>
	//	1. Save messages as pdf
	//	2. Upload pdf to lighthouse
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
	if err := bitcoin.ModifyAndBroadcastClosingTx(os.Getenv("recipientAddress"), closingTxHex, rpcClient, metadata); err != nil {
		logger.LogE("Error modifying and broadcasting transaction: ", err)
	}

}
