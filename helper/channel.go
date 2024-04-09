package helper

import (
	"MasterThesis/logger"
	lightning "github.com/chainpoint/lightning-go"
	"github.com/lightningnetwork/lnd/lnrpc"
)

func CreateChannel(lnd *lightning.LightningClient, p2p string, satoshi int64) lnrpc.Lightning_OpenChannelClient {
	channelClient, err := lnd.CreateChannel(p2p, satoshi)
	if err != nil {
		logger.LogE("Lightning node create channel error: ", err)
	}
	return channelClient
}
