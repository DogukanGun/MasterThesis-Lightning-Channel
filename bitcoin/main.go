package bitcoin

import (
	"bytes"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"log"
	"os"
)

func ModifyAndBroadcastClosingTx(recipientAddress string, rawTxHex string, client *rpcclient.Client, metadata string) error {
	rawTx := []byte(rawTxHex)

	msgTx := wire.NewMsgTx(wire.TxVersion)
	if err := msgTx.Deserialize(bytes.NewReader(rawTx)); err != nil {
		return err
	}

	// Decode recipient address
	recipientAddr, err := btcutil.DecodeAddress(recipientAddress, &chaincfg.RegressionNetParams)
	if err != nil {
		return err
	}
	pkScript, err := txscript.PayToAddrScript(recipientAddr)
	if err != nil {
		return err
	}

	// Add OP_RETURN output with metadata
	opReturnScript, err := txscript.NullDataScript([]byte(metadata))
	if err != nil {
		return err
	}
	txOut := wire.NewTxOut(0, opReturnScript)
	msgTx.AddTxOut(txOut)

	// Decode private key
	privKey, err := btcutil.DecodeWIF(os.Getenv("privKeyWIF"))
	if err != nil {
		log.Fatalf("Error decoding private key: %v", err)
	}

	// Sign the transaction
	for i, txIn := range msgTx.TxIn {
		sigScript, err := txscript.SignatureScript(msgTx, i, pkScript, txscript.SigHashAll, privKey.PrivKey, true)
		if err != nil {
			return err
		}
		txIn.SignatureScript = sigScript
	}

	txHash, err := client.SendRawTransaction(msgTx, true)
	if err != nil {
		return err
	}

	log.Printf("Transaction sent successfully with txid: %s", txHash)
	return nil
}
