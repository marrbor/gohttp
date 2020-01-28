package gohttp_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marrbor/gohttp"
	"github.com/stretchr/testify/assert"
)

type testRequest struct {
	ID     int64    `json:"id"`
	Name   string   `json:"name"`
	Params []string `json:"params"`
}

var testReq = testRequest{
	ID:     123,
	Name:   "abcdefg",
	Params: []string{"hij", "klmn", "opqr", "stu", "vw", "xyz"},
}

func TestRequestJSONToParams(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data testRequest
		err := gohttp.RequestJSONToParams(r, &data)
		assert.NoError(t, err)
		assert.EqualValues(t, testReq.ID, data.ID)
		assert.EqualValues(t, testReq.Name, data.Name)
		assert.EqualValues(t, testReq.Params, data.Params)
		gohttp.ResponseOK(w)
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	body, err := json.Marshal(testReq)
	assert.NoError(t, err)

	r, err := http.Post(ts.URL, "text/json", bytes.NewReader(body))
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusOK, r.StatusCode)
}

func TestRequestJSONToParams2(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data testRequest
		err := gohttp.RequestJSONToParams(r, &data)
		assert.Error(t, err)
		gohttp.BadRequest(w, err)
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	// send illegal type to generate unmarshal error.
	invalidReq := struct{ ID string }{ID: "id"}
	body, err := json.Marshal(invalidReq)
	assert.NoError(t, err)

	r, err := http.Post(ts.URL, "text/json", bytes.NewReader(body))
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusBadRequest, r.StatusCode)
	assert.EqualValues(t, "400 Bad Request", r.Status)
}

func TestGenRequest(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.EqualValues(t, r.Method, http.MethodGet)
		assert.EqualValues(t, "", r.Header.Get("Content-Type"))
		gohttp.ResponseOK(w)
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	url := ts.URL
	req, err := gohttp.GenRequest(gohttp.HttpMethods.GET, url, nil)
	assert.NoError(t, err)
	c := new(http.Client)
	res, err := c.Do(req)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
}

func TestGenRequest2(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.EqualValues(t, http.MethodPost, r.Method)
		assert.EqualValues(t, "application/json", r.Header.Get("Content-Type"))
		var data testRequest
		err := gohttp.RequestJSONToParams(r, &data)
		assert.NoError(t, err)
		assert.EqualValues(t, testReq.ID, data.ID)
		assert.EqualValues(t, testReq.Name, data.Name)
		assert.EqualValues(t, testReq.Params, data.Params)
		gohttp.ResponseOK(w)
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	url := ts.URL
	req, err := gohttp.GenRequest(gohttp.HttpMethods.POST, url, &testReq)
	assert.NoError(t, err)
	c := new(http.Client)
	res, err := c.Do(req)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
}

func TestGenRequest3(t *testing.T) {
	// give channel as parameter to generate JSON marshall error.
	req, err := gohttp.GenRequest(gohttp.HttpMethods.PUT, "http://localhost:8080", make(chan int))
	assert.Nil(t, req)
	assert.Error(t, err)
}
