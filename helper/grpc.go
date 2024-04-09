package helper

import (
	"MasterThesis/logger"
	"crypto/tls"
	"crypto/x509"
	"github.com/lightningnetwork/lnd/lncfg"
	"github.com/lightningnetwork/lnd/macaroons"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"gopkg.in/macaroon.v2"
	"io/ioutil"
	"math"
	"os"
)

var MACAROONOPTION grpc.CallOption

func GrpcSetup(port string, certFileLocation string, macaroonFileLocation string) *grpc.ClientConn {
	/*usr, err := user.Current()
	if err != nil {
		logger.LogE(err)
	}
	homeDir := usr.HomeDir
	lndDir := fmt.Sprintf("%s%s", homeDir, os.Getenv("POLAR_PATH"))*/
	// SSL credentials setup
	//var serverName string
	//certFileLocation := os.Getenv("POLAR_TLS")
	f, err := os.ReadFile(certFileLocation)

	p := x509.NewCertPool()
	p.AppendCertsFromPEM(f)
	tlsConfig := &tls.Config{
		RootCAs: p,
	}
	creds := credentials.NewTLS(tlsConfig)
	//creds, err := credentials.NewClientTLSFromFile(certFileLocation, serverName)
	if err != nil {
		logger.LogE(err)
	}
	// Macaroon setup/Users/dogukangundogan/
	///data/chain/bitcoin/testnet/admin.macaroon
	//macaroonFileLocation := os.Getenv("POLAR_MACAROON")
	macaroonMap := map[string]string{"macaroon": macaroonFileLocation}
	macaroonMetadata := metadata.New(macaroonMap)
	MACAROONOPTION = grpc.Header(&macaroonMetadata)
	macBytes, err := ioutil.ReadFile(macaroonFileLocation)

	if err != nil {
		logger.LogE(err)
	}
	mac := &macaroon.Macaroon{}
	if err = mac.UnmarshalBinary(macBytes); err != nil {
		logger.LogE(err.Error())
	}
	macConstraints := []macaroons.Constraint{
		macaroons.TimeoutConstraint(60),
	}
	// Apply constraints to the macaroon.
	constrainedMac, err := macaroons.AddConstraints(mac, macConstraints...)
	if err != nil {
		logger.LogE(err)
	}
	// Now we append the macaroon credentials to the dial options.
	cred, _ := macaroons.NewMacaroonCredential(constrainedMac)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	//port := os.Getenv("POLAR_PORT")
	opts = append(opts, grpc.WithPerRPCCredentials(cred))
	genericDialer := lncfg.ClientAddressDialer(port)
	opts = append(opts, grpc.WithContextDialer(genericDialer))
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt)))
	opts = append(opts, grpc.WithInitialConnWindowSize(math.MaxInt32))
	opts = append(opts, grpc.WithInitialWindowSize(math.MaxInt32))
	opts = append(opts, grpc.WithMaxHeaderListSize(math.MaxInt32))
	conn, err := grpc.Dial(os.Getenv("POLAR"), opts...)

	if err != nil {
		logger.LogE(err)
	}
	logger.LogI("Connection done")
	return conn
}
