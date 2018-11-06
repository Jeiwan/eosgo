package eosgo

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

// GET is an HTTP request helper function
func GET(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		fallthrough
	case 201:
		fallthrough
	case 202:

	default:
		return nil, errors.New(string(respBody))
	}

	return respBody, nil
}

// POST is an HTTP request helper function
func POST(url string, body []byte) ([]byte, error) {
	reqBody := bytes.NewBuffer(body)
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		fallthrough
	case 201:
		fallthrough
	case 202:

	default:
		return nil, errors.New(string(respBody))
	}

	return respBody, nil
}
