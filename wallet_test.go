package eosgo_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jeiwan/eosgo"

	"github.com/Jeiwan/eosgo/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignTransaction(t *testing.T) {
	m := http.NewServeMux()
	m.HandleFunc("/v1/wallet/sign_transaction", func(w http.ResponseWriter, r *http.Request) {
		var reqBody []interface{}

		err := json.NewDecoder(r.Body).Decode(&reqBody)
		require.Nil(t, err)

		unsignedTxJSON, _ := json.Marshal(reqBody[0].(map[string]interface{}))

		var unsignedTx types.RawTransaction
		json.Unmarshal(unsignedTxJSON, &unsignedTx)

		unsignedTx.Signatures = []string{"SIGNATURE"}

		resp, err := json.Marshal(unsignedTx)
		if err != nil {
			fmt.Println(err)
			require.Nil(t, err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(resp))
	})

	s := httptest.NewServer(m)
	defer s.Close()

	eos := eosgo.New(eosgo.EOSConfig{
		NodeosURL: "not used",
		KeosURL:   s.URL,
	})

	tx := &types.RawTransaction{
		Actions: []types.RawAction{
			types.RawAction{
				Account: "eosio",
				Name:    "test",
				Authorization: []types.Authorization{
					types.Authorization{Actor: "test", Permission: "active"},
				},
				Data: json.RawMessage(`"deadbeef"`),
			},
		},
		DelaySec: 123,
	}

	err := eos.SignTransaction(tx, "PUBLIC_KEY")
	if err != nil {
		fmt.Println(err)
		require.Nil(t, err)
	}

	require.Len(t, tx.Signatures, 1)
	assert.Equal(t, "SIGNATURE", tx.Signatures[0])
}
