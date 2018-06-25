package eosgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Jeiwan/eosgo/types"
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
func (a ChainAPI) GetInfo() (*types.Info, error) {
	resp, err := http.Get(a.PathPrefix + "/get_info")
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	var info types.Info
	if err = json.Unmarshal(respBody, &info); err != nil {
		return nil, err
	}

	return &info, nil
}
