package types

import (
	"encoding/json"
	"github.com/golang/glog"
	"io"
	"io/ioutil"
)

func BscERC20ProphecyClaimToJsonString(e *BscERC20ProphecyClaim) string {
	bytes, err := json.Marshal(e)
	if err != nil {
		glog.Errorf("BscERC20ProphecyClaimToJsonString err:%+v", err)
		return ""
	}
	return string(bytes)
}

func JsonToBscERC20ProphecyClaim(r io.Reader) (*BscERC20ProphecyClaim, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		glog.Errorf("read bscerc20prophecyclaim request body err: %+v", err)
		return nil, err
	}

	epc := BscERC20ProphecyClaim{}
	err = json.Unmarshal(body, &epc)
	if err != nil {
		glog.Errorf("convert string to BscERC20ProphecyClaim err: %+v", err)
		return nil, err
	}

	return &epc, nil
}

func EthERC20ProphecyClaimToJsonString(e *EthERC20ProphecyClaim) string {
	bytes, err := json.Marshal(e)
	if err != nil {
		glog.Errorf("EthERC20ProphecyClaimToJsonString err:%+v", err)
		return ""
	}
	return string(bytes)
}

func JsonToEthERC20ProphecyClaim(r io.Reader) (*EthERC20ProphecyClaim, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		glog.Errorf("read etherc20prophecyclaim request body err: %+v", err)
		return nil, err
	}

	epc := EthERC20ProphecyClaim{}
	err = json.Unmarshal(body, &epc)
	if err != nil {
		glog.Errorf("convert string to EthProphecyClaim err: %+v", err)
		return nil, err
	}

	return &epc, nil
}







func BscERC721ProphecyClaimToJsonString(e *BscERC721ProphecyClaim) string {
	bytes, err := json.Marshal(e)
	if err != nil {
		glog.Errorf("BscERC721ProphecyClaimToJsonString err:%+v", err)
		return ""
	}
	return string(bytes)
}

func JsonToBscERC721ProphecyClaim(r io.Reader) (*BscERC721ProphecyClaim, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		glog.Errorf("read bscerc20prophecyclaim request body err: %+v", err)
		return nil, err
	}

	epc := BscERC721ProphecyClaim{}
	err = json.Unmarshal(body, &epc)
	if err != nil {
		glog.Errorf("convert string to BscERC721ProphecyClaim err: %+v", err)
		return nil, err
	}

	return &epc, nil
}

func EthERC721ProphecyClaimToJsonString(e *EthERC721ProphecyClaim) string {
	bytes, err := json.Marshal(e)
	if err != nil {
		glog.Errorf("EthERC721ProphecyClaimToJsonString err:%+v", err)
		return ""
	}
	return string(bytes)
}

func JsonToEthERC721ProphecyClaim(r io.Reader) (*EthERC721ProphecyClaim, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		glog.Errorf("read etherc20prophecyclaim request body err: %+v", err)
		return nil, err
	}

	epc := EthERC721ProphecyClaim{}
	err = json.Unmarshal(body, &epc)
	if err != nil {
		glog.Errorf("convert string to EthProphecyClaim err: %+v", err)
		return nil, err
	}

	return &epc, nil
}