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
	defer r.Body.Close()
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

// GenRequest generate an HTTP request for send. Body have to have `json` tag to be marshal to JSON strings.
func GenRequest(method HTTPMethod, url string, body interface{}) (*http.Request, error) {
	var buf io.Reader = nil
	if body != nil {
		bj, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(bj)
	}

	req, err := http.NewRequest(method.String(), url, buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	return req, nil
}
