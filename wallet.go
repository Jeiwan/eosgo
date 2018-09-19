package eosgo

import (
	"encoding/json"

	"github.com/Jeiwan/eosgo/types"
)

// SignTransaction signs a transaction
func (eos EOS) SignTransaction(tx *types.RawTransaction, publicKey string) error {
	reqBody := []interface{}{
		tx,
		[]string{publicKey},
		"",
	}

	reqBodyData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	respBody, err := POST(eos.Config.KeosURL+"/v1/wallet/sign_transaction", reqBodyData)
	if err != nil {
		return err
	}

	var signedTx types.RawTransaction
	if err = json.Unmarshal(respBody, &signedTx); err != nil {
		return err
	}

	*tx = signedTx

	return nil
}
