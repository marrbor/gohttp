package gohttp_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marrbor/gohttp"
	"github.com/stretchr/testify/assert"
)

type testResponse struct {
	ID     int64    `json:"id"`
	Name   string   `json:"name"`
	Params []string `json:"params"`
}

var res = testResponse{
	ID:     123,
	Name:   "abcdefg",
	Params: []string{"hij", "klmn", "opqr", "stu", "vw", "xyz"},
}

func TestBadRequest(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gohttp.BadRequest(w, fmt.Errorf("bad request"))
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusBadRequest, r.StatusCode)
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, "bad request\n", string(body))
}

func TestForbidden(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gohttp.Forbidden(w, fmt.Errorf("forbidden"))
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusForbidden, r.StatusCode)
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, "forbidden\n", string(body))
}

func TestInternalServerError(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gohttp.InternalServerError(w, fmt.Errorf("internal server error"))
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, r.StatusCode)
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, "internal server error\n", string(body))
}

func TestJSONResponse(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := gohttp.JSONResponse(w, &res)
		assert.NoError(t, err)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusOK, r.StatusCode)

	var resp testResponse
	err = gohttp.ResponseJSONToParams(r, &resp)
	assert.NoError(t, err)
	assert.EqualValues(t, res.ID, resp.ID)
	assert.EqualValues(t, res.Name, resp.Name)
	assert.EqualValues(t, res.Params, resp.Params)
}

func TestJSONResponse2(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := gohttp.JSONResponse(w, nil)
		assert.NoError(t, err)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusOK, r.StatusCode)
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, 0, len(body))
}

func TestUnauthorized(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gohttp.Unauthorized(w, fmt.Errorf("unauthorized"))
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, r.StatusCode)
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, "unauthorized\n", string(body))
}

func TestMethodNotAllowed(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gohttp.MethodNotAllowed(w, fmt.Errorf("method not allowed"))
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusMethodNotAllowed, r.StatusCode)
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, "method not allowed\n", string(body))
}

func TestNotFound(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gohttp.NotFound(w, fmt.Errorf("not found"))
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusNotFound, r.StatusCode)
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, "not found\n", string(body))

}

func TestResponseOK(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gohttp.ResponseOK(w)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusOK, r.StatusCode)
}

func TestResponseJSONToParams(t *testing.T) {
	// only test error path since normal path is tested in TestJSONResponse.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := gohttp.JSONResponse(w, &res)
		assert.NoError(t, err)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusOK, r.StatusCode)

	// give invalid structure.
	var x struct{ ID string }
	err = gohttp.ResponseJSONToParams(r, &x)
	assert.Error(t, err)
}
