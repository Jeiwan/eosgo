package eosgo_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Jeiwan/eosgo"
	"github.com/Jeiwan/eosgo/types"
	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/v1/chain/get_info", func(w http.ResponseWriter, r *http.Request) {
		resp := `
		{
			"server_version": "db031363",
			"chain_id": "aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906",
			"head_block_num": 2556258,
			"last_irreversible_block_num": 2555927,
			"last_irreversible_block_id": "0027001704d9cabf1cfc1e624cd76a06b1f7a37ca4e3acf89e18bc585d07a82e",
			"head_block_id": "002701629e16aa1b3c5c07038dc264a3c042dfaa0ed1103e3bb29c749ded8cc1",
			"head_block_time": "2018-06-25T11:16:27.000",
			"head_block_producer": "eosgenblockp",
			"virtual_block_cpu_limit": 200000000,
			"virtual_block_net_limit": 1048576000,
			"block_cpu_limit": 192272,
			"block_net_limit": 1048240
		}
		`

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, resp)
	})

	s := httptest.NewServer(m)
	defer s.Close()

	api := eosgo.NewChainAPI(s.URL)
	resp, err := api.GetInfo()

	assert.Nil(t, err)

	assert.Equal(t, "db031363", resp.ServerVersion)
	assert.Equal(t, "aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906", resp.ChainID)
	assert.Equal(t, 2556258, resp.HeadBlockNum)
	assert.Equal(t, 2555927, resp.LastIrreversibleBlockNum)
	assert.Equal(t, "0027001704d9cabf1cfc1e624cd76a06b1f7a37ca4e3acf89e18bc585d07a82e", resp.LastIrreversibleBlockID)
	assert.Equal(t, "002701629e16aa1b3c5c07038dc264a3c042dfaa0ed1103e3bb29c749ded8cc1", resp.HeadBlockID)
	assert.Equal(t, types.NewTime(time.Date(2018, time.June, 25, 11, 16, 27, 0, time.UTC)), resp.HeadBlockTime)
	assert.Equal(t, "eosgenblockp", resp.HeadBlockProducer)
	assert.Equal(t, 200000000, resp.VirtualBlockCPULimit)
	assert.Equal(t, 1048576000, resp.VirtualBlockNetLimit)
	assert.Equal(t, 192272, resp.BlockCPULimit)
	assert.Equal(t, 1048240, resp.BlockNetLimit)
}

func TestGetEmptyBlock(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/v1/chain/get_block", func(w http.ResponseWriter, r *http.Request) {
		resp := `
		{
			"timestamp": "2018-06-08T08:08:08.500",
			"producer": "dan",
			"confirmed": 1,
			"previous": "0000000000000000000000000000000000000000000000000000000000000000",
			"transaction_mroot": "0000000000000000000000000000000000000000000000000000000000000000",
			"action_mroot": "aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906",
			"schedule_version": 0,
			"new_producers": null,
			"header_extensions": [],
			"producer_signature": "SIG_K1_111111111111111111111111111111111111111111111111111111111111111116uk5ne",
			"transactions": [],
			"block_extensions": [],
			"id": "00000001405147477ab2f5f51cda427b638191c66d2c59aa392d5c2c98076cb0",
			"block_num": 1,
			"ref_block_prefix": 4126519930
		}
		`

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, resp)
	})

	s := httptest.NewServer(m)
	defer s.Close()

	api := eosgo.NewChainAPI(s.URL)
	resp, err := api.GetBlockByNumber(1)

	assert.Nil(t, err)
	assert.Equal(t, types.NewTime(time.Date(2018, 6, 8, 8, 8, 8, 500000000, time.UTC)), resp.Timestmap)
	assert.Equal(t, "dan", resp.Producer)
	assert.Equal(t, 1, resp.Confirmed)
	assert.Equal(t, "0000000000000000000000000000000000000000000000000000000000000000", resp.Previous)
	assert.Equal(t, "0000000000000000000000000000000000000000000000000000000000000000", resp.TransactionMroot)
	assert.Equal(t, "aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906", resp.ActionMroot)
	assert.Equal(t, 0, resp.ScheduleVersion)
	// assert.Equal(t, nil, resp.NewProducers)
	// assert.Equal(t, nil, resp.HeadExtensions)
	assert.Equal(t, "SIG_K1_111111111111111111111111111111111111111111111111111111111111111116uk5ne", resp.ProducerSignature)
	assert.Equal(t, []types.Transaction{}, resp.Transactions)
	// assert.Equal(t, nil, resp.BlockExtensions)
	assert.Equal(t, "00000001405147477ab2f5f51cda427b638191c66d2c59aa392d5c2c98076cb0", resp.ID)
	assert.Equal(t, 1, resp.BlockNum)
	assert.Equal(t, 4126519930, resp.RefBlockPrefix)
}
