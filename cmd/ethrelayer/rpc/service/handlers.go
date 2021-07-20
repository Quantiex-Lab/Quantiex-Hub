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
	util.AddHandler("POST", util.RELAY_BINANCE_ERC20PROPHECYCLAIM, HandleRelayERC20ProphecyClaim)
	util.AddHandler("POST", util.RELAY_BINANCE_ERC721PROPHECYCLAIM, HandleRelayERC721ProphecyClaim)

	port := rpc.GetConfig().Port
	util.SetPort(port)

	go startMessageLoop(s)
	go util.Start()
}

func HandleRelayERC20ProphecyClaim(c echo.Context) error {
	bscProphecyClaim, err := types.JsonToBscERC20ProphecyClaim(c.Request().Body)
	if err != nil {
		glog.Errorf("jsonToClaimType err: %+v", err)
		return err
	}

	RelayERC20ProphecyClaim(bscProphecyClaim)

	return nil
}

func HandleRelayERC721ProphecyClaim(c echo.Context) error {
	bscProphecyClaim, err := types.JsonToBscERC721ProphecyClaim(c.Request().Body)
	if err != nil {
		glog.Errorf("jsonToClaimType err: %+v", err)
		return err
	}

	RelayERC721ProphecyClaim(bscProphecyClaim)

	return nil
}
