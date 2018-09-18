package eosgo

import (
	"encoding/json"

	"github.com/Jeiwan/eosgo/types"
)

/*
	TODO:
      get_block_header_state
      get_account
      get_code
      get_abi
      get_raw_code_and_abi
      get_table_rows
      get_currency_balance
      get_currency_stats
      get_producers
      get_producer_schedule
      get_scheduled_transactions
      abi_json_to_bin
      abi_bin_to_json
      get_required_keys
      get_transaction_id
      push_block
      push_transaction
      push_transactions
*/

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
