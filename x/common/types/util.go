package types

import (
	"encoding/json"
	"github.com/golang/glog"
	"io"
	"io/ioutil"
)

func BscProphecyClaimToJsonString(e *BscProphecyClaim) string {
	bytes, err := json.Marshal(e)
	if err != nil {
		glog.Errorf("BscProphecyClaimToJsonString err:%+v", err)
		return ""
	}
	return string(bytes)
}

func JsonToBscProphecyClaim(r io.Reader) (*BscProphecyClaim, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		glog.Errorf("read bscprophecyclaim request body err: %+v", err)
		return nil, err
	}

	epc := BscProphecyClaim{}
	err = json.Unmarshal(body, &epc)
	if err != nil {
		glog.Errorf("convert string to BscProphecyClaim err: %+v", err)
		return nil, err
	}

	return &epc, nil
}

func EthProphecyClaimToJsonString(e *EthProphecyClaim) string {
	bytes, err := json.Marshal(e)
	if err != nil {
		glog.Errorf("EthProphecyClaimToJsonString err:%+v", err)
		return ""
	}
	return string(bytes)
}

func JsonToEthProphecyClaim(r io.Reader) (*EthProphecyClaim, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		glog.Errorf("read ethprophecyclaim request body err: %+v", err)
		return nil, err
	}

	epc := EthProphecyClaim{}
	err = json.Unmarshal(body, &epc)
	if err != nil {
		glog.Errorf("convert string to EthProphecyClaim err: %+v", err)
		return nil, err
	}

	return &epc, nil
}
