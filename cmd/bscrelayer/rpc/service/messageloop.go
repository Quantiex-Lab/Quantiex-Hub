package service

import (
	"github.com/Quantiex-Hub/cmd/bscrelayer/relayer"
	"github.com/Quantiex-Hub/cmd/bscrelayer/txs"
	"github.com/Quantiex-Hub/cmd/bscrelayer/types"
	xcommon "github.com/Quantiex-Hub/x/common/types"
	"github.com/golang/glog"
)

type msgType int

const (
	relayERC20ProphecyClaim msgType = iota
	relayERC721ProphecyClaim
)

var mgr *messageMgr

type message struct {
	op    msgType
	param interface{}
}

type messageMgr struct {
	msgChan chan *message
	sub     *relayer.BinanceSub
}

func startMessageLoop(s *relayer.BinanceSub) {
	mgr = &messageMgr{
		msgChan: make(chan *message, 10000),
		sub:     s,
	}
	mgr.messageLoop()
}

func (m *messageMgr) messageLoop() {
	for {
		select {
		case msg, open := <-m.msgChan:
			if !open {
				glog.Info("chan closed, message loop exit")
				return
			}
			switch msg.op {
			case relayERC20ProphecyClaim:
				glog.Info("message loop msg: relayERC20ProphecyClaim")
				if msg.param == nil {
					glog.Info("relayERC20ProphecyClaim param nil")
					continue
				}

				pc, ok := msg.param.(*xcommon.EthERC20ProphecyClaim)
				if !ok {
					glog.Info("addPubChan param err")
					continue
				}
				m.handleRelayERC20ProphecyClaimMsg(pc)
			case relayERC721ProphecyClaim:
				glog.Info("message loop msg: relayERC721ProphecyClaim")
				if msg.param == nil {
					glog.Info("relayERC721ProphecyClaim param nil")
					continue
				}

				pc, ok := msg.param.(*xcommon.EthERC721ProphecyClaim)
				if !ok {
					glog.Info("addPubChan param err")
					continue
				}
				m.handleRelayERC721ProphecyClaimMsg(pc)
			}
		}
	}
}

func (m *messageMgr) isLoopExit() bool {
	return m.msgChan == nil
}

func RelayERC20ProphecyClaim(prophecyClaim *xcommon.EthERC20ProphecyClaim) {
	if mgr.isLoopExit() {
		glog.Errorf("channel is close, relayERC20ProphecyClaim msg not implement")
		return
	}
	glog.Infof("RelayProphecyClaim prophecyClaim is:%+v", prophecyClaim)

	mgr.msgChan <- &message{op: relayERC20ProphecyClaim, param: prophecyClaim}
	return
}

func RelayERC721ProphecyClaim(prophecyClaim *xcommon.EthERC721ProphecyClaim) {
	if mgr.isLoopExit() {
		glog.Errorf("channel is close, EthERC721ProphecyClaim msg not implement")
		return
	}
	glog.Infof("RelayProphecyClaim prophecyClaim is:%+v", prophecyClaim)

	mgr.msgChan <- &message{op: relayERC721ProphecyClaim, param: prophecyClaim}
	return
}

func (m *messageMgr) handleRelayERC20ProphecyClaimMsg(prophecyClaim *xcommon.EthERC20ProphecyClaim) {
	if m.isLoopExit() {
		glog.Errorf("channel is close, handleRelayProphecyClaimMsg msg not implement")
		return
	}

	var claimType types.Event
	if prophecyClaim.ClaimType == xcommon.LockText {
		claimType = types.MsgLock
	} else if prophecyClaim.ClaimType == xcommon.BurnText {
		claimType = types.MsgBurn
	}

	err := txs.RelayERC20ProphecyClaimToBinance(m.sub.BscProvider, m.sub.RegistryContractAddress,
		claimType, *prophecyClaim, m.sub.PrivateKey)
	if err != nil {
		glog.Errorf("handleRelayProphecyClaimMsg err:%+v", err)
		return
	}

	return
}

func (m *messageMgr) handleRelayERC721ProphecyClaimMsg(prophecyClaim *xcommon.EthERC721ProphecyClaim) {
	if m.isLoopExit() {
		glog.Errorf("channel is close, handleRelayProphecyClaimMsg msg not implement")
		return
	}

	var claimType types.Event
	if prophecyClaim.ClaimType == xcommon.LockText {
		claimType = types.MsgLock
	} else if prophecyClaim.ClaimType == xcommon.BurnText {
		claimType = types.MsgBurn
	}
	err := txs.RelayERC721ProphecyClaimToBinance(m.sub.BscProvider, m.sub.RegistryContractAddress,
		claimType, *prophecyClaim, m.sub.PrivateKey)
	if err != nil {
		glog.Errorf("handleRelayProphecyClaimMsg err:%+v", err)
		return
	}
	return
}
