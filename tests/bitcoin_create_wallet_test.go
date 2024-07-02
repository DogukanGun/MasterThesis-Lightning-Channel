package tests

import (
	"MasterThesis/logger"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// bcrt1q0r5vx0pm0r4hg2klqcd7jhpzm2dyf8stqjq7vh
func TestBitcoinCreateWallet(t *testing.T) {
	t.Setenv("BitcoinHost", "127.0.0.1:18449")
	t.Setenv("BitcoinUser", "polaruser")
	t.Setenv("BitcoinPass", "polarpass")
	t.Setenv("PrevTransaction", "35490db69f58588a675df834f84ac3239a4fe0cc392dc82c355f6729dbd8509f")
	t.Setenv("TransactionHex", "020000000001057282218adf21ab1d5e047c4325fc0793bb1cc1ad4dbb62a38268b8f47e69a7050100000000fdffffffac14d33b558dcc0ae9899e911c84b4fe5c278a59cb0111804f5cca9d074df16f0000000000fdffffffdecd1c23b62beacfeddf2bfabc13259bec301870f375b3b93f5a72d683099f920000000000fdffffffc36256f3c1ea337a300855be71a96070f916b867569449215a1f90906b3d178b0000000000fdffffff4de847bc7148fc08f73d3c5cfa3cab4429fd982b1830bb8ec2ade98c3cc3db1e0000000000fdffffff0180841e00000000001600140d21ee6223da4fdd3b5db385023271cad56f10c30247304402205a40ffd7702efe08b687ea9ee3e0f5e6e048c662dd321bed58e839b86807b2d00220296611648360932339a0b89df7158f5faaecff08f132faecd67525d7930892bc012102bdf25c7f00892e393f0cb0a960a1498a0827c957c0d84989ea9e51c2375dec6702463043021f7c5028c2f55c67c8503c1e2c3f2dc1118c534d7070ed465d36781e38021a83022060c46045863aaf552af032e6b4ea47a6f4f187a58eef285d61ca582f87f0a58e012102293e99551df219d6dd50f30071052c9c01a3a351b69dbc682b09cead2164c2f60247304402206fa5bb47e042dc7ff2f192d8b120ba3391516e1a412a80ad4697cf73d1010e4602207db56893c5efb3d3cb01b6af54e26836e7506afa8c471754cfeb9a980d7509b0012102293e99551df219d6dd50f30071052c9c01a3a351b69dbc682b09cead2164c2f602473044022030c175a76bc7ecf7eee4b44843e2668ff5de32ba8cf23f4975397a57cc8dad1b022029bbfff0ac1e6dc9338d9893aeec856002ca2148cfaa1675885eb8c8f57b3f20012102293e99551df219d6dd50f30071052c9c01a3a351b69dbc682b09cead2164c2f6024730440220441cb7e4c01a36e0cee7bade58e85b12e1014282f6a9faf64fd01313d5e4febc0220111a6b599d7308aadec51d0e0fc7ed3a4f532ab663ba72cf668b3b696fd22133012102293e99551df219d6dd50f30071052c9c01a3a351b69dbc682b09cead2164c2f68b100000")
	rpcClient, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         os.Getenv("BitcoinHost"),
		User:         os.Getenv("BitcoinUser"),
		Pass:         os.Getenv("BitcoinPass"),
		HTTPPostMode: true,
		DisableTLS:   true,
	}, nil)
	res, err := rpcClient.CreateWallet("test2")
	if err != nil {
		logger.LogE(err)
		return
	}
	walletInfo, err := rpcClient.GetWalletInfo()
	if err != nil {
		logger.LogE(err)
		return
	}
	logger.LogI(walletInfo)
	if err != nil {
		logger.LogE(err)
		return
	}
	assert.NotEqual(t, "", res.Name)
}
