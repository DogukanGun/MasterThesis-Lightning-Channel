package helper

import (
	"MasterThesis/logger"
	"fmt"
	lightning "github.com/chainpoint/lightning-go"
	"os"
	"os/user"
	"time"
)

func ConnectNode(lndClient *lightning.LightningClient) {
	err := lndClient.WaitForConnection(5 * time.Minute)
	if err != nil {
		logger.LogE(err.Error())
	}
	err = lndClient.Unlocker()
	if err != nil {
		logger.LogE(err.Error())
	}
}

func GetDefaultNode() lightning.LightningClient {
	return lightning.LightningClient{
		TlsPath:        "/Users/dogukangundogan/.polar/networks/1/volumes/lnd/dave/tls.cert",
		MacPath:        "/Users/dogukangundogan/.polar/networks/1/volumes/lnd/dave/data/chain/bitcoin/regtest/admin.macaroon",
		ServerHostPort: "127.0.0.1:10004",
		LndLogLevel:    "error",
		MinConfs:       3,
		Testnet:        true,
		HashPrice:      int64(2), //price to charge for issuing LSAT
	}
}

func GetLndNode() lightning.LightningClient {
	usr, err := user.Current()
	if err != nil {
		logger.LogE(err)
	}
	homeDir := usr.HomeDir
	lndDir := fmt.Sprintf("%s/Library/Application Support/Lnd", homeDir)

	// SSL credentials setup
	//var serverName string
	certFileLocation := lndDir + os.Getenv("LND_TLS")
	return lightning.LightningClient{
		TlsPath:        certFileLocation,
		MacPath:        lndDir + os.Getenv("LND_MACAROON"),
		ServerHostPort: "127.0.0.1:8080",
		LndLogLevel:    "error",
		MinConfs:       3,
		Testnet:        true,
		HashPrice:      int64(2), //price to charge for issuing LSAT
	}
}
