package gohttp_test

import (
	"github.com/marrbor/gohttp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var methods = []string{"POST", "GET", "PUT", "DELETE"}
var headers = []string{"X-TEST-HEAD", "Content-Type"}
var origin = "google.com"

func TestAddCORSHeader(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gohttp.AddCORSHeader(w, origin)
		gohttp.ResponseOK(w)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusOK, r.StatusCode)
	assert.EqualValues(t, origin, r.Header.Get("Access-Control-Allow-Origin"))
}

func TestPreflightForCORS(t *testing.T) {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.False(t, gohttp.PreflightForCORS(w, r, &headers, &methods, origin))
		gohttp.AddCORSHeader(w, "google.com")
		gohttp.ResponseOK(w)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusOK, r.StatusCode)
	assert.EqualValues(t, "google.com", r.Header.Get("Access-Control-Allow-Origin"))
}

func TestPreflightForCORS2(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, gohttp.PreflightForCORS(w, r, nil, nil, origin))
		gohttp.AddCORSHeader(w, "google.com")
		gohttp.ResponseOK(w)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	req, err := http.NewRequest(http.MethodOptions, ts.URL, nil)
	assert.NoError(t, err)
	req.Header.Add("Access-Control-Allow-Origin", origin)
	req.Header.Add("Access-Control-Request-Headers", strings.Join(headers, ", "))
	req.Header.Add("Access-Control-Request-Method", strings.Join(methods, ", "))
	client := new(http.Client)
	r, err := client.Do(req)

	assert.EqualValues(t, http.StatusOK, r.StatusCode)
	assert.EqualValues(t, "google.com", r.Header.Get("Access-Control-Allow-Origin"))
	assert.EqualValues(t, strings.Join(headers, ", "), r.Header.Get("Access-Control-Allow-Headers"))
	assert.EqualValues(t, strings.Join(methods, ", "), r.Header.Get("Access-Control-Allow-Methods", ))
}

func TestPreflightForCORS3(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, gohttp.PreflightForCORS(w, r, &headers, &methods, origin))
		gohttp.AddCORSHeader(w, "google.com")
		gohttp.ResponseOK(w)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	req, err := http.NewRequest(http.MethodOptions, ts.URL, nil)
	assert.NoError(t, err)
	client := new(http.Client)
	r, err := client.Do(req)
	assert.NotNil(t, r)
	assert.EqualValues(t, http.StatusOK, r.StatusCode)
	assert.EqualValues(t, "google.com", r.Header.Get("Access-Control-Allow-Origin"))
	assert.EqualValues(t, strings.Join(headers, ", "), r.Header.Get("Access-Control-Allow-Headers"))
	assert.EqualValues(t, strings.Join(methods, ", "), r.Header.Get("Access-Control-Allow-Methods", ))

}
