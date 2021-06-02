package service

import (
	"github.com/Quantiex-Hub/cmd/ethrelayer/relayer"
	"github.com/Quantiex-Hub/cmd/ethrelayer/rpc"
	"github.com/Quantiex-Hub/cmd/util"
	"github.com/Quantiex-Hub/x/common/types"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

func StartHttpServer(s *relayer.EthereumSub) {
	util.AddHandler("POST", util.BINANCE_RELAYPROPHECYCLAIM, HandleRelayProphecyClaim)

	port := rpc.GetConfig().Port
	util.SetPort(port)

	go startMessageLoop(s)
	go util.Start()
}

func HandleRelayProphecyClaim(c echo.Context) error {
	bscProphecyClaim, err := types.JsonToBscProphecyClaim(c.Request().Body)
	if err != nil {
		glog.Errorf("jsonToClaimType err: %+v", err)
		return err
	}

	RelayProphecyClaim(bscProphecyClaim)

	return nil
}
