package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Node allows to communicate with an EOS node via its HTTP API
type Node struct {
	APIEndpoint string
}

func NewNode(endpointURL string) *Node {
	return &Node{
		APIEndpoint: endpointURL,
	}
}

// GetTableRows allows to read table content of certain contract
func (n Node) GetTableRows(code, scope, table string) (*TableRows, error) {
	reqMap := map[string]interface{}{
		"code":  code,
		"scope": scope,
		"table": table,
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
