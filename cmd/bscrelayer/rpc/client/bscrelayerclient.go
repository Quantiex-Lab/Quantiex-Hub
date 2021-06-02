package client

import (
	"github.com/Quantiex-Hub/cmd/bscrelayer/rpc"
	"github.com/Quantiex-Hub/cmd/util"
	"github.com/Quantiex-Hub/x/common/types"
	"github.com/golang/glog"
)

func SendProphecyClaimToEthereum(claim types.BscProphecyClaim) (string, error) {
	cfg := rpc.GetConfig()
	URL := cfg.Url

	URL += util.BINANCE_RELAYPROPHECYCLAIM
	connect := util.GetHTTPClient()
	claimString := types.BscProphecyClaimToJsonString(&claim)
	body, _, err := util.HttpGet("POST", URL, claimString, nil, nil, connect)
	if err != nil {
		glog.Error(err)
		return "", err
	}

	return body, nil
}
