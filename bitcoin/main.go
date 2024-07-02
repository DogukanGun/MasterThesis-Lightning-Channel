package bitcoin

import (
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

// bitcoin-cli -regtest getaddressinfo bcrt1qk3hwlvjfju5xvdqrhnn7f9m37aumh7k9snzzxr
// bitcoin-cli getnewaddress

func PublishMetadata(prevTxHashStr string, prevOutIndex uint32, metadata string, client *rpcclient.Client) error {
	// Create a new empty transaction
	msgTx := wire.NewMsgTx(wire.TxVersion)

	// Convert the previous transaction hash from string to chainhash.Hash
	prevTxHash, err := chainhash.NewHashFromStr(prevTxHashStr)
	if err != nil {
		return err
	}

	// Define the previous outpoint
	prevOutPoint := wire.NewOutPoint(prevTxHash, prevOutIndex)

	// Create the transaction input with the witness data
	txIn := wire.NewTxIn(prevOutPoint, nil, nil)
	msgTx.AddTxIn(txIn)

	// Define the recipient address
	receiver, err := btcutil.DecodeAddress("bcrt1q0r5vx0pm0r4hg2klqcd7jhpzm2dyf8stqjq7vh", &chaincfg.RegressionNetParams)
	if err != nil {
		return err
	}

	// Create the output script for the recipient address
	receiverScript, err := txscript.PayToAddrScript(receiver)
	if err != nil {
		return err
	}

	// Add the transaction output (send 1,000 satoshis to the recipient)
	txOut := wire.NewTxOut(1000, receiverScript)
	msgTx.AddTxOut(txOut)

	// Add the OP_RETURN output with metadata
	metadataScript, err := txscript.NullDataScript([]byte(metadata))
	if err != nil {
		return err
	}
	opReturnTxOut := wire.NewTxOut(0, metadataScript)
	msgTx.AddTxOut(opReturnTxOut)
	// Sign the transaction (assuming you have the private key)
	signedTx, err := signTransaction(msgTx, client)
	if err != nil {
		return err
	}

	// Broadcast the transaction
	txHash, err := client.SendRawTransaction(signedTx, true)
	if err != nil {
		return err
	}

	// Convert chainhash.Hash to string
	txHashStr := txHash.String()
	fmt.Printf("Transaction sent successfully with txid: %s\n", txHashStr)
	return nil
}

func signTransaction(msgTx *wire.MsgTx, client *rpcclient.Client) (*wire.MsgTx, error) {
	// Use the wallet to sign the transaction
	signedTx, complete, err := client.SignRawTransactionWithWallet(msgTx)
	if err != nil {
		return nil, err
	}
	if !complete {
		return nil, errors.New("transaction signing incomplete")
	}

	return signedTx, nil
}
