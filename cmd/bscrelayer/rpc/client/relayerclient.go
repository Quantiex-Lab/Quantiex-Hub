package client

import (
	"github.com/Quantiex-Hub/cmd/bscrelayer/rpc"
	"github.com/Quantiex-Hub/cmd/util"
	"github.com/Quantiex-Hub/x/common/types"
	"github.com/golang/glog"
)

func SendERC20ProphecyClaimToEthereum(claim types.BscERC20ProphecyClaim) (string, error) {
	cfg := rpc.GetConfig()
	URL := cfg.Url

	URL += util.RELAY_BINANCE_ERC20PROPHECYCLAIM
	connect := util.GetHTTPClient()
	claimString := types.BscERC20ProphecyClaimToJsonString(&claim)
	body, _, err := util.HttpGet("POST", URL, claimString, nil, nil, connect)
	if err != nil {
		glog.Error(err)
		return "", err
	}

	return body, nil
}

func SendERC721ProphecyClaimToEthereum(claim types.BscERC721ProphecyClaim) (string, error) {
	cfg := rpc.GetConfig()
	URL := cfg.Url

	URL += util.RELAY_BINANCE_ERC721PROPHECYCLAIM
	connect := util.GetHTTPClient()
	claimString := types.BscERC721ProphecyClaimToJsonString(&claim)
	body, _, err := util.HttpGet("POST", URL, claimString, nil, nil, connect)
	if err != nil {
		glog.Error(err)
		return "", err
	}

	return body, nil
}
