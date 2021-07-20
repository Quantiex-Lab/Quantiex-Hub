package client

import (
	"github.com/Quantiex-Hub/cmd/ethrelayer/rpc"
	"github.com/Quantiex-Hub/cmd/util"
	"github.com/Quantiex-Hub/x/common/types"
	"github.com/golang/glog"
)

func SendERC20ProphecyClaimToBinance(claim types.EthERC20ProphecyClaim) (string, error) {
	cfg := rpc.GetConfig()
	URL := cfg.Url

	URL += util.RELAY_ETHEREUM_ERC20PROPHECYCLAIM
	connect := util.GetHTTPClient()
	claimString := types.EthERC20ProphecyClaimToJsonString(&claim)
	body, _, err := util.HttpGet("POST", URL, claimString, nil, nil, connect)
	if err != nil {
		glog.Error(err)
		return "", err
	}

	return body, nil
}

func SendERC721ProphecyClaimToBinance(claim types.EthERC721ProphecyClaim) (string, error) {
	cfg := rpc.GetConfig()
	URL := cfg.Url

	URL += util.RELAY_ETHEREUM_ERC721PROPHECYCLAIM
	connect := util.GetHTTPClient()
	claimString := types.EthERC721ProphecyClaimToJsonString(&claim)
	body, _, err := util.HttpGet("POST", URL, claimString, nil, nil, connect)
	if err != nil {
		glog.Error(err)
		return "", err
	}

	return body, nil
}