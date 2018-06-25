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
