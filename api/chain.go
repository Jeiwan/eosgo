package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ABIJSONtoBin converts JSON representation of ABI to binary
func (n Node) ABIJSONtoBin(contractAccount, action string, jsonArgs []string) (string, error) {
	reqMap := map[string]interface{}{
		"code":   contractAccount,
		"action": action,
		"args":   jsonArgs,
	}

	reqData, err := json.Marshal(reqMap)
	if err != nil {
		return "", nil
	}

	reqBody := bytes.NewBuffer(reqData)
	url := n.APIEndpoint + "/v1/chain/abi_json_to_bin"

	client := http.Client{}
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var respBody map[string]interface{}
	err = json.Unmarshal(respBodyData, &respBody)
	if err != nil {
		return "", err
	}

	binargs, ok := respBody["binargs"]
	if !ok {
		return "", fmt.Errorf("reseponse doesn't contain 'binargs': %s", respBodyData)
	}

	return binargs.(string), nil
}

// GetTableRows allows to read table content of certain contract
func (n Node) GetTableRows(code, scope, table string) (*TableRows, error) {
	reqMap := map[string]interface{}{
		"code":  code,
		"scope": scope,
		"table": table,
		"json":  true,
	}

	reqData, err := json.Marshal(reqMap)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(reqData)
	url := n.APIEndpoint + "/v1/chain/get_table_rows"

	client := http.Client{}
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tableRows TableRows
	err = json.Unmarshal(respBody, &tableRows)
	if err != nil {
		return nil, err
	}

	return &tableRows, nil

}

// TableRows response chain/get_table_rows
type TableRows struct {
	Rows []map[string]interface{} `json:"rows"`
	More bool                     `json:"more"`
}
