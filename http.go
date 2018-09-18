package eosgo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GET is an HTTP request helper function
func GET(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed request: %s", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return respBody, nil
}

// POST is an HTTP request helper function
func POST(url string, body []byte) ([]byte, error) {
	reqBody := bytes.NewBuffer(body)
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed request; %s", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return respBody, nil
}
