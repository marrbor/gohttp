package gohttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
