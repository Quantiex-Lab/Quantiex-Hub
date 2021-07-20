package service

import (
	"github.com/Quantiex-Hub/cmd/bscrelayer/relayer"
	"github.com/Quantiex-Hub/cmd/bscrelayer/rpc"
	"github.com/Quantiex-Hub/cmd/util"
	"github.com/Quantiex-Hub/x/common/types"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

func StartHttpServer(s *relayer.BinanceSub) {
	util.AddHandler("POST", util.RELAY_ETHEREUM_ERC20PROPHECYCLAIM, HandleRelayERC20ProphecyClaim)
	util.AddHandler("POST", util.RELAY_ETHEREUM_ERC721PROPHECYCLAIM, HandleRelayERC721ProphecyClaim)

	port := rpc.GetConfig().Port
	util.SetPort(port)

	go startMessageLoop(s)
	go util.Start()
}

func HandleRelayERC20ProphecyClaim(c echo.Context) error {
	ethProphecyClaim, err := types.JsonToEthERC20ProphecyClaim(c.Request().Body)
	if err != nil {
		glog.Errorf("jsonToClaimType err: %+v", err)
		return err
	}

	RelayERC20ProphecyClaim(ethProphecyClaim)

	return nil
}

func HandleRelayERC721ProphecyClaim(c echo.Context) error {
	ethProphecyClaim, err := types.JsonToEthERC721ProphecyClaim(c.Request().Body)
	if err != nil {
		glog.Errorf("jsonToClaimType err: %+v", err)
		return err
	}

	RelayERC721ProphecyClaim(ethProphecyClaim)

	return nil
}
