package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	eostypes "github.com/Jeiwan/eosgo/types"
)

// SignTransaction signs a transaction
func (n Node) SignTransaction(tx *eostypes.Transaction, publicKey string, chainID int) error {
	reqBodyData, err := json.Marshal(tx)
	if err != nil {
		return err
	}

	reqBody := bytes.NewBuffer(reqBodyData)
	url := n.APIEndpoint + "/v1/wallet/sign_transaction"

	client := http.Client{}
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var signedTx map[string]interface{}
	err = json.Unmarshal(respBody, &signedTx)
	if err != nil {
		return err
	}

	fmt.Println(signedTx)

	return nil
}
