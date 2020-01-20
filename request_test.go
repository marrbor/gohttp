package gohttp_test

import (
	"bytes"
	"encoding/json"
	"github.com/marrbor/gohttp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testRequest struct {
	ID     int64    `json:"id"`
	Name   string   `json:"name"`
	Params []string `json:"params"`
}

var req = testRequest{
	ID:     123,
	Name:   "abcdefg",
	Params: []string{"hij", "klmn", "opqr", "stu", "vw", "xyz"},
}

func TestRequestJSONToParams(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data testRequest
		err := gohttp.RequestJSONToParams(r, &data)
		assert.NoError(t, err)
		assert.EqualValues(t, req.ID, data.ID)
		assert.EqualValues(t, req.Name, data.Name)
		for i, s := range req.Params {
			assert.EqualValues(t, req.Params[i], s)
		}
		gohttp.ResponseOK(w)
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	body, err := json.Marshal(req)
	assert.NoError(t, err)

	r, err := http.Post(ts.URL, "text/json", bytes.NewReader(body))
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusOK, r.StatusCode)
}
