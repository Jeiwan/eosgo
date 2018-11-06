package eosgo

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Jeiwan/eosgo/types"
)

/*
	TODO:
      abi_bin_to_json
      get_abi
      get_account
      get_block_header_state
      get_code
      get_producer_schedule
      get_producers
      get_raw_code_and_abi
      get_required_keys
      get_scheduled_transactions
      get_transaction_id
      push_block
      push_transactions
*/

// GetAccount returns account information by its name
func (eos EOS) GetAccount(name string) (*GetAccountResponse, error) {
	reqBody := map[string]interface{}{
		"account_name": name,
	}

	reqBodyData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	respBody, err := POST(eos.Config.NodeosURL+"/v1/chain/get_account", reqBodyData)
	if err != nil {
		return nil, err
	}

	var account GetAccountResponse
	if err = json.Unmarshal(respBody, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

// GetCurrencyBalance returns account currency balance
func (eos EOS) GetCurrencyBalance(code, account, symbol string) ([]string, error) {
	reqBody := map[string]interface{}{
		"code":    code,
		"account": account,
		"symbol":  symbol,
	}

	reqBodyData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	respBody, err := POST(eos.Config.NodeosURL+"/v1/chain/get_currency_balance", reqBodyData)
	if err != nil {
		return nil, err
	}

	var balanceResponse []string
	if err = json.Unmarshal(respBody, &balanceResponse); err != nil {
		return nil, err
	}

	return balanceResponse, nil
}

// GetCurrencyStats returns currency stats
func (eos EOS) GetCurrencyStats(code, symbol string) (map[string]CurrencyStats, error) {
	reqBody := map[string]interface{}{
		"code":   code,
		"symbol": symbol,
	}

	reqBodyData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	respBody, err := POST(eos.Config.NodeosURL+"/v1/chain/get_currency_stats", reqBodyData)
	if err != nil {
		return nil, err
	}

	var statsResponse map[string]CurrencyStats
	if err = json.Unmarshal(respBody, &statsResponse); err != nil {
		return nil, err
	}

	return statsResponse, nil
}

// GetInfo returns blockchain information
func (eos EOS) GetInfo() (*types.Info, error) {
	respBody, err := GET(eos.Config.NodeosURL + "/v1/chain/get_info")
	if err != nil {
		return nil, err
	}

	var info types.Info
	if err = json.Unmarshal(respBody, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

// GetBlockByNumber retrieves a block by its number
func (eos EOS) GetBlockByNumber(number int) (*types.Block, error) {
	reqBody := map[string]int{
		"block_num_or_id": number,
	}
	reqBodyData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	respBody, err := POST(eos.Config.NodeosURL+"/v1/chain/get_block", reqBodyData)
	if err != nil {
		return nil, err
	}

	var block types.Block
	if err = json.Unmarshal(respBody, &block); err != nil {
		return nil, err
	}

	return &block, nil
}

// GetBlockByID retrieves a block by its ID
func (eos EOS) GetBlockByID(id string) (*types.Block, error) {
	reqBody := map[string]string{
		"block_num_or_id": id,
	}
	reqBodyData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	respBody, err := POST(eos.Config.NodeosURL+"/v1/chain/get_block", reqBodyData)
	if err != nil {
		return nil, err
	}

	var block types.Block
	if err = json.Unmarshal(respBody, &block); err != nil {
		return nil, err
	}

	return &block, nil
}

// GetTableRows reads contract's table and returns its rows
func (eos EOS) GetTableRows(contract, scope, table string) (*GetTableRowsResponse, error) {
	reqBody := map[string]interface{}{
		"code":  contract,
		"scope": scope,
		"table": table,
		"limit": 1000,
		"json":  true,
	}

	reqBodyData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	respBody, err := POST(eos.Config.NodeosURL+"/v1/chain/get_table_rows", reqBodyData)
	if err != nil {
		return nil, err
	}

	var response GetTableRowsResponse
	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// PushTransaction sends transaction to the blockchain
func (eos EOS) PushTransaction(tx *types.RawTransaction) error {
	reqBody := map[string]interface{}{
		"transaction":       tx,
		"signatures":        tx.Signatures,
		"context_free_data": tx.ContextFreeData,
		"compression":       "none",
	}

	reqBodyData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	_, err = POST(eos.Config.NodeosURL+"/v1/chain/push_transaction", reqBodyData)
	if err != nil {
		return err
	}

	return nil
}

// ABIJSONtoBin converts JSON representation of ABI to binary
func (eos EOS) ABIJSONtoBin(contractAccount, action string, jsonArgs []interface{}) (string, error) {
	reqMap := map[string]interface{}{
		"code":   contractAccount,
		"action": action,
		"args":   jsonArgs,
	}

	reqData, err := json.Marshal(reqMap)
	if err != nil {
		return "", err
	}

	respBody, err := POST(eos.Config.NodeosURL+"/v1/chain/abi_json_to_bin", reqData)
	if err != nil {
		return "", err
	}

	var resp map[string]interface{}
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return "", err
	}

	binargs, ok := resp["binargs"]
	if !ok {
		return "", fmt.Errorf("reseponse doesn't contain 'binargs': %s", respBody)
	}

	return binargs.(string), nil
}

// GetAccountResponse represents chain/get_account response
type GetAccountResponse struct {
	AccountName       string     `bson:"account_name" json:"account_name"`
	HeadBlockNum      int        `bson:"head_block_num" json:"head_block_num"`
	HeadBlockTime     types.Time `bson:"head_block_time" json:"head_block_time"`
	Privileged        bool       `bson:"privileged" json:"privileged"`
	LastCodeUpdate    types.Time `bson:"last_code_update" json:"last_code_update"`
	Created           types.Time `bson:"created" json:"created"`
	CoreLiquidBalance string     `bson:"core_liquid_balance" json:"core_liquid_balance"`
	RAMQuota          numOrStr   `bson:"ram_quota" json:"ram_quota"`
	NetWeight         numOrStr   `bson:"net_weight" json:"net_weight"`
	CPUWeight         numOrStr   `bson:"cpu_weight" json:"cpu_weight"`
	RAMUsage          numOrStr   `bson:"ram_usage" json:"ram_usage"`
	CPULimit          struct {
		Available numOrStr `bson:"available" json:"available"`
		Max       numOrStr `bson:"max" json:"max"`
		Used      numOrStr `bson:"used" json:"used"`
	} `bson:"cpu_limit" json:"cpu_limit"`
	NetLimit struct {
		Available numOrStr `bson:"available" json:"available"`
		Max       numOrStr `bson:"max" json:"max"`
		Used      numOrStr `bson:"used" json:"used"`
	} `bson:"net_limit" json:"net_limit"`
	Permissions []struct {
		PermName     string `bson:"perm_name" json:"perm_name"`
		Parent       string `bson:"parent" json:"parent"`
		RequiredAuth struct {
			Threshold int `bson:"threshold" json:"threshold"`
			Keys      []struct {
				Key    string `bson:"key" json:"key"`
				Weight int    `bson:"weight" json:"weight"`
			} `bson:"keys" json:"keys"`
			Accounts []struct {
				Permission types.Authorization `bson:"permission" json:"permission"`
				Weight     int                 `bson:"weight" json:"weight"`
			}
			// Waits ... TODO
		} `bson:"required_auth" json:"required_auth"`
	} `bson:"permissions" json:"permissions"`
	TotalResources struct {
		Owner     string   `bson:"owner" json:"owner"`
		NetWeight string   `bson:"net_weight" json:"net_weight"`
		CPUWeight string   `bson:"cpu_weight" json:"cpu_weight"`
		RAMBytes  numOrStr `bson:"ram_bytes" json:"ram_bytes"`
	} `bson:"total_resources" json:"total_resources"`
	SelfDelegatedBandwidth struct {
		From      string `bson:"from" json:"from"`
		To        string `bson:"to" json:"to"`
		NetWeight string `bson:"net_weight" json:"net_weight"`
		CPUWeight string `bson:"cpu_weight" json:"cpu_weight"`
	} `bson:"self_delegated_bandwidth" json:"self_delegated_bandwidth"`
	// RefundRequest  TODO
	VoterInfo struct {
		Owner             string   `bson:"owner" json:"owner"`
		Proxy             string   `bson:"proxy" json:"proxy"`
		Producers         []string `bson:"producers" json:"producers"`
		Staked            int      `bson:"staked" json:"staked"`
		TotalVoteWeight   string   `bson:"last_vote_weight" json:"last_vote_weight"`
		ProxiedVoteWeight string   `bson:"proxied_vote_weight" json:"proxied_vote_weight"`
		IsProxy           int      `bson:"is_proxy" json:"is_proxy"`
	} `bson:"voter_info" json:"voter_info"`
}

type numOrStr int

func (mi *numOrStr) UnmarshalJSON(data []byte) error {
	type mirror numOrStr
	var m mirror

	err := json.Unmarshal(data, &m)
	if err != nil {
		if err.Error() == "json: cannot unmarshal string into Go value of type main.mirror" {
			var str string
			err = json.Unmarshal(data, &str)
			if err != nil {
				return err
			}

			i, err := strconv.Atoi(str)
			if err != nil {
				return err
			}

			m = mirror(i)
		} else {
			return err
		}
	}

	*mi = numOrStr(m)

	return nil
}

// GetTableRowsResponse represents a response from chain/get_table_row
type GetTableRowsResponse struct {
	Rows []map[string]interface{}
	More bool
}

// CurrencyStats represents currency stats as returned by GetCurrencyStats
type CurrencyStats struct {
	Supply    string `json:"supply"`
	MaxSupply string `json:"max_supply"`
	Issuer    string `json:"issuer"`
}
