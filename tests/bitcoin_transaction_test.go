package tests

import (
	"MasterThesis/bitcoin"
	"MasterThesis/logger"
	"github.com/btcsuite/btcd/rpcclient"
	"os"
	"testing"
)

func TestBitcoinTransaction(t *testing.T) {
	t.Setenv("BitcoinHost", "127.0.0.1:18443")
	t.Setenv("BitcoinUser", "polaruser")
	t.Setenv("BitcoinPass", "polarpass")
	rpcClient, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         os.Getenv("BitcoinHost"),
		User:         os.Getenv("BitcoinUser"),
		Pass:         os.Getenv("BitcoinPass"),
		HTTPPostMode: true,
		DisableTLS:   true,
	}, nil)
	if err != nil {
		logger.LogE(err)
	}
	metadata := "This is a test"
	previousTranx := ""
	if err := bitcoin.PublishMetadata(previousTranx, metadata, rpcClient); err != nil {
		logger.LogE("Error modifying and broadcasting transaction: ", err.Error())
	}
}
