package eosgo

import (
	"encoding/json"

	"github.com/Jeiwan/eosgo/types"
)

/*
	TODO:
		create
		create_key
		get_public_keys
		import_key
		list_keys
		list_wallets
		lock
		lock_all
		open
		remove_key
		set_timeout
		sign_digest
		sign_transaction
		unlock
*/

// SignTransaction signs a transaction
func (eos EOS) SignTransaction(tx *types.RawTransaction, publicKey, chainID string) error {
	reqBody := []interface{}{
		tx,
		[]string{publicKey},
		chainID,
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
