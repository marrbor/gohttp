package gohttp

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

////// Receive request (for server side program)

// RequestJSONToParams convert request JSON body to given structure.
func RequestJSONToParams(r *http.Request, params interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &params); err != nil {
		return err
	}
	return nil
}

////// Send request (for client side program)

// GenRequest generate an HTTP request for send. 'data' can be put multiple data, but only first one data is taken. Currently, data will be marshal to JSON strings.
func GenRequest(method, url string, data ... interface{}) (*http.Request, error) {
	var body io.Reader = nil
	if data != nil {
		bj, err := json.Marshal(data[0])
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(bj)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	return req, nil
}
