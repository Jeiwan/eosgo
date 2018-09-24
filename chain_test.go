package eosgo_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Jeiwan/eosgo"
	eostypes "github.com/Jeiwan/eosgo/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCurrencyBalance(t *testing.T) {
	t.Run("ok", func(tt *testing.T) {
		m := http.NewServeMux()
		m.HandleFunc("/v1/chain/get_currency_balance", func(w http.ResponseWriter, r *http.Request) {
			resp := `["3.1337 EOS"]`

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, resp)
		})

		s := httptest.NewServer(m)
		defer s.Close()

		eos := eosgo.New(eosgo.EOSConfig{NodeosURL: s.URL})
		resp, err := eos.GetCurrencyBalance("eos.token", "test", "EOS")
		require.Nil(t, err)

		assert.Equal(tt, []string{"3.1337 EOS"}, resp)
	})
}

func TestGetInfo(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/v1/chain/get_info", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadFile("./fixtures/info.json")

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(resp))
	})

	s := httptest.NewServer(m)
	defer s.Close()

	eos := eosgo.New(eosgo.EOSConfig{NodeosURL: s.URL})
	resp, err := eos.GetInfo()

	require.Nil(t, err)

	assert.Equal(t, "db031363", resp.ServerVersion)
	assert.Equal(t, "aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906", resp.ChainID)
	assert.Equal(t, 2556258, resp.HeadBlockNum)
	assert.Equal(t, 2555927, resp.LastIrreversibleBlockNum)
	assert.Equal(t, "0027001704d9cabf1cfc1e624cd76a06b1f7a37ca4e3acf89e18bc585d07a82e", resp.LastIrreversibleBlockID)
	assert.Equal(t, "002701629e16aa1b3c5c07038dc264a3c042dfaa0ed1103e3bb29c749ded8cc1", resp.HeadBlockID)
	assert.Equal(t, eostypes.NewTime(time.Date(2018, time.June, 25, 11, 16, 27, 0, time.UTC)), resp.HeadBlockTime)
	assert.Equal(t, "eosgenblockp", resp.HeadBlockProducer)
	assert.Equal(t, 200000000, int(resp.VirtualBlockCPULimit.(float64)))
	assert.Equal(t, 1048576000, int(resp.VirtualBlockNetLimit.(float64)))
	assert.Equal(t, 192272, resp.BlockCPULimit)
	assert.Equal(t, 1048240, resp.BlockNetLimit)
}

func TestGetEmptyBlock(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/v1/chain/get_block", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadFile("./fixtures/empty_block.json")

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(resp))
	})

	s := httptest.NewServer(m)
	defer s.Close()

	eos := eosgo.New(eosgo.EOSConfig{NodeosURL: s.URL})
	resp, err := eos.GetBlockByNumber(1)

	assert.Nil(t, err)
	assert.Equal(t, eostypes.NewTime(time.Date(2018, 6, 8, 8, 8, 8, 500000000, time.UTC)), resp.Timestamp)
	assert.Equal(t, "dan", resp.Producer)
	assert.Equal(t, 1, resp.Confirmed)
	assert.Equal(t, "0000000000000000000000000000000000000000000000000000000000000000", string(resp.Previous))
	assert.Equal(t, "0000000000000000000000000000000000000000000000000000000000000000", string(resp.TransactionMroot))
	assert.Equal(t, "aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906", string(resp.ActionMroot))
	assert.Equal(t, 0, resp.ScheduleVersion)
	// assert.Equal(t, nil, resp.NewProducers)
	// assert.Equal(t, nil, resp.HeadExtensions)
	assert.Equal(t, "SIG_K1_111111111111111111111111111111111111111111111111111111111111111116uk5ne", resp.ProducerSignature)
	assert.Empty(t, resp.Transactions)
	// assert.Equal(t, nil, resp.BlockExtensions)
	assert.Equal(t, "00000001405147477ab2f5f51cda427b638191c66d2c59aa392d5c2c98076cb0", string(resp.ID))
	assert.Equal(t, 1, resp.BlockNum)
	assert.Equal(t, 4126519930, resp.RefBlockPrefix)
}

func TestGetFullBlock(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/v1/chain/get_block", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadFile("./fixtures/full_block.json")

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(resp))
	})

	s := httptest.NewServer(m)
	defer s.Close()

	eos := eosgo.New(eosgo.EOSConfig{NodeosURL: s.URL})
	resp, err := eos.GetBlockByNumber(1)

	assert.Nil(t, err)
	assert.Equal(t, eostypes.NewTime(time.Date(2018, 6, 9, 12, 6, 33, 0, time.UTC)), resp.Timestamp)
	assert.Equal(t, "eosio", resp.Producer)
	assert.Equal(t, 0, resp.Confirmed)
	assert.Equal(t, "000003e7e53c1e971717ebeae28d30e6cf8d1d4c8f246f978592f4c6df27d1bc", string(resp.Previous))
	assert.Equal(t, "ef919ff6ed63578c7632e7029ac4eb6ec1fceeaf7f5629e9b36bded9ed4f3c49", string(resp.TransactionMroot))
	assert.Equal(t, "d91ddc84af3b36dc80375ca988c4c5421e16aa0622f11343669ffdb2be800027", string(resp.ActionMroot))
	assert.Equal(t, 0, resp.ScheduleVersion)
	// assert.Equal(t, nil, resp.NewProducers)
	// assert.Equal(t, nil, resp.HeadExtensions)
	assert.Equal(t, "SIG_K1_JxqacVaf1dLYqgqUwcnzfgnzDVqnsu9Z5TK3Jk1EqQnnw3vAPYwVHdwdvyqA8oiUYDybpq8GSkneKAaqeEwevP9GJBLLkQ", resp.ProducerSignature)

	txH := resp.Transactions[0]
	assert.Equal(t, "executed", txH.Status)
	assert.Equal(t, 100395, txH.CPUUsageUs)
	assert.Equal(t, 1298, txH.NetUsageWords)

	trx := txH.Trx
	assert.Equal(t, "1bc395276f4bdde15a7992e50e61938457673e861d9480b51762b6e4457e5b79", string(trx.ID))
	assert.Equal(t, []string{"SIG_K1_KAzq4AycYweK2tLKpZYmiMaf5xrJLZVinrPSLSiPN9E5eca9cgmBFTJcvYcyY2A1TkKDg7LXxf7TgbTunTSLdxfHPiY71J"}, trx.Signatures)
	assert.Equal(t, "none", trx.Compression)
	assert.Empty(t, trx.PackedContextFreeData)
	assert.Equal(t, []interface{}{}, trx.ContextFreeData)
	assert.Equal(t, "w00t", string(trx.PackedTrx))

	tx := trx.Transaction
	assert.Equal(t, eostypes.NewTime(time.Date(2018, 6, 9, 13, 6, 32, 0, time.UTC)), tx.Expiration)
	assert.Equal(t, 997, tx.RefBlockNum)
	assert.Equal(t, 2927439535, tx.RefBlockPrefix)
	assert.Equal(t, 0, tx.MaxNetUsagWords)
	assert.Equal(t, 0, tx.MaxCPUUsageMs)
	assert.Equal(t, 0, tx.DelaySec)
	assert.Equal(t, []eostypes.Action{}, tx.ContextFreeActions)

	a := tx.Actions[0]
	assert.Equal(t, "eosio", a.Account)
	assert.Equal(t, "newaccount", a.Name)

	auth := a.Authorization[0]
	assert.Equal(t, "eosio", auth.Actor)
	assert.Equal(t, "active", auth.Permission)

	data := a.Data
	assert.Equal(t, "eosio", data["creator"])
	assert.Equal(t, "ha4tqmjxgege", data["name"])

	owner := data["owner"].(map[string]interface{})
	assert.Equal(t, 1.0, owner["threshold"])

	key := owner["keys"].([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "EOS8Y4gqJ9ZTQZf41cozxHm15MtxTXCMkcjS2TmeEjfJeK3ZU1ooi", key["key"])
	assert.Equal(t, 1.0, key["weight"])
	assert.Empty(t, owner["accounts"])
	assert.Empty(t, owner["waits"])

	active := data["active"].(map[string]interface{})
	assert.Equal(t, 1.0, active["threshold"])
	key = active["keys"].([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "EOS8Y4gqJ9ZTQZf41cozxHm15MtxTXCMkcjS2TmeEjfJeK3ZU1ooi", key["key"])
	assert.Equal(t, 1.0, key["weight"])
	assert.Empty(t, owner["accounts"])
	assert.Empty(t, owner["waits"])

	assert.Equal(t, "deadbeef0", string(*a.HexData))

	a = tx.Actions[1]
	assert.Equal(t, "eosio", a.Account)
	assert.Equal(t, "buyrambytes", a.Name)

	auth = a.Authorization[0]
	assert.Equal(t, "eosio", auth.Actor)
	assert.Equal(t, "active", auth.Permission)

	data = a.Data
	assert.Equal(t, "eosio", data["payer"])
	assert.Equal(t, "ha4tqmjxgege", data["receiver"])
	assert.Equal(t, 8192.0, data["bytes"])

	assert.Equal(t, "0000000000ea3055a09862fd499b896900200000", string(*a.HexData))

	a = tx.Actions[2]
	assert.Equal(t, "eosio", a.Account)
	assert.Equal(t, "delegatebw", a.Name)

	auth = a.Authorization[0]
	assert.Equal(t, "eosio", auth.Actor)
	assert.Equal(t, "active", auth.Permission)

	data = a.Data
	assert.Equal(t, "eosio", data["from"])
	assert.Equal(t, "ha4tqmjxgege", data["receiver"])
	assert.Equal(t, "4.0000 EOS", data["stake_net_quantity"])
	assert.Equal(t, "4.0000 EOS", data["stake_cpu_quantity"])
	assert.Equal(t, 1.0, data["transfer"])

	assert.Equal(t, "deadbeef1", string(*a.HexData))

	a = tx.Actions[3]
	assert.Equal(t, "eosio.token", a.Account)
	assert.Equal(t, "transfer", a.Name)

	auth = a.Authorization[0]
	assert.Equal(t, "eosio", auth.Actor)
	assert.Equal(t, "active", auth.Permission)

	data = a.Data
	assert.Equal(t, "eosio", data["from"])
	assert.Equal(t, "ha4tqmjxgege", data["to"])
	assert.Equal(t, "2.0000 EOS", data["quantity"])
	assert.Equal(t, "init", data["memo"])

	assert.Equal(t, "deadbeef2", string(*a.HexData))

	// assert.Equal(t, nil, resp.BlockExtensions)
	assert.Equal(t, "000003e8ddc03486114e6a1c764f0f78dab559ed18519802b4399a8d89e48264", string(resp.ID))
	assert.Equal(t, 1000, resp.BlockNum)
	assert.Equal(t, 476728849, resp.RefBlockPrefix)
}

func TestPushTransaction(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/v1/chain/push_transaction", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "")
	})

	s := httptest.NewServer(m)
	defer s.Close()

	eos := eosgo.New(eosgo.EOSConfig{NodeosURL: s.URL})

	tx := &eostypes.RawTransaction{
		Actions: []eostypes.RawAction{
			eostypes.RawAction{
				Account: "eosio",
				Name:    "test",
				Authorization: []eostypes.Authorization{
					eostypes.Authorization{Actor: "test", Permission: "active"},
				},
				Data: json.RawMessage(`"deadbeef"`),
			},
		},
		DelaySec: 123,
	}

	err := eos.PushTransaction(tx)
	require.Nil(t, err)
}
