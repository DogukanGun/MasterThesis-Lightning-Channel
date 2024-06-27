package bitcoin

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"log"
)

func PublishMetadata(prevTxHashStr string, metadata string, client *rpcclient.Client) error {
	// Create a new empty transaction
	msgTx := wire.NewMsgTx(wire.TxVersion)

	prevTxHash, _ := chainhash.NewHashFromStr(prevTxHashStr)
	prevOutPoint := wire.NewOutPoint(prevTxHash, 0) // 0 is the index of the output in the previous transaction
	txIn := wire.NewTxIn(prevOutPoint, nil, nil)
	msgTx.AddTxIn(txIn)

	// Add OP_RETURN output with metadata
	opReturnScript, err := txscript.NullDataScript([]byte(metadata))
	if err != nil {
		return err
	}
	txOut := wire.NewTxOut(0, opReturnScript)
	msgTx.AddTxOut(txOut)

	// Send the transaction
	txHash, err := client.SendRawTransaction(msgTx, true)
	if err != nil {
		return err
	}

	log.Printf("Transaction sent successfully with txid: %s", txHash)
	return nil
}
