package channel

import (
	"MasterThesis/logger"
	"context"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
)

func OpenChannel(destinationPubKey string, grpcConn *grpc.ClientConn) {
	lncli := lnrpc.NewLightningClient(grpcConn)
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
