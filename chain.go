package eosgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ChainAPI ...
type ChainAPI struct {
	PathPrefix string
}

// NewChainAPI ...
func NewChainAPI(url string) *ChainAPI {
	return &ChainAPI{PathPrefix: fmt.Sprintf("%s/v1/chain", url)}
}

// GetInfo ...
func (a ChainAPI) GetInfo() (*Info, error) {
	resp, err := http.Get(a.PathPrefix + "/get_info")
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	var info Info
	if err = json.Unmarshal(respBody, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

type getBlockByNumberReq struct {
	Number int `json:"block_num_or_id"`
}

// GetBlockByNumber ...
func (a ChainAPI) GetBlockByNumber(number int) (*Block, error) {
	req := getBlockByNumberReq{Number: number}

	reqRaw, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(reqRaw)
	resp, err := http.Post(a.PathPrefix+"/get_block", "application/json", reqBody)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	var block Block
	if err = json.Unmarshal(respBody, &block); err != nil {
		return nil, err
	}

	return &block, nil
}
