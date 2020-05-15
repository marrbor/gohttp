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

func TestNotImplementedError(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gohttp.NotImplementedError(w, fmt.Errorf("not implemented"))
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusNotImplemented, r.StatusCode)
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, "not implemented\n", string(body))
}

func TestBadGatewayError(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gohttp.BadGatewayError(w, fmt.Errorf("bad gateway"))
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusBadGateway, r.StatusCode)
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, "bad gateway\n", string(body))
}

func TestServiceUnavailableError(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gohttp.ServiceUnavailableError(w, fmt.Errorf("service unavailable"))
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusServiceUnavailable, r.StatusCode)
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, "service unavailable\n", string(body))
}

func TestGatewayTimeoutError(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gohttp.GatewayTimeoutError(w, fmt.Errorf("gateway timeout"))
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()
	r, err := http.Get(ts.URL)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusGatewayTimeout, r.StatusCode)
	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.EqualValues(t, "gateway timeout\n", string(body))
}

func isTestSubroutine(t *testing.T, base int, f func(r *http.Response) bool) {
	r := new(http.Response)

	// out of range patterns for every function
	r.StatusCode = -1
	assert.EqualValues(t, false, f(r))
	r.StatusCode = 0
	assert.EqualValues(t, false, f(r))
	r.StatusCode = 1
	assert.EqualValues(t, false, f(r))

	r.StatusCode = base - 1
	assert.EqualValues(t, false, f(r))
	r.StatusCode = base
	assert.EqualValues(t, true, f(r))
	r.StatusCode = base + 1
	assert.EqualValues(t, true, f(r))
	r.StatusCode = base + 99
	assert.EqualValues(t, true, f(r))
	r.StatusCode = base + 100
	assert.EqualValues(t, false, f(r))
}

func TestIs(t *testing.T) {
	isTestSubroutine(t, 100, gohttp.IsInformational)
	isTestSubroutine(t, 200, gohttp.IsSuccessful)
	isTestSubroutine(t, 300, gohttp.IsRedirection)
	isTestSubroutine(t, 400, gohttp.IsClientError)
	isTestSubroutine(t, 500, gohttp.IsServerError)

	r := new(http.Response)
	r.StatusCode = 404
	assert.True(t, gohttp.IsNotFound(r))
	r.StatusCode = 300
	assert.False(t, gohttp.IsNotFound(r))
}

func TestResponseToError(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := gohttp.JSONResponse(w, &res)
		assert.NoError(t, err)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	r, _ := http.Get(ts.URL)
	r.StatusCode = 404
	r.Status = "404 Not found"
	err := gohttp.ResponseToError(r, nil, nil)
	assert.EqualError(t, err, "404 404 Not found:\n {\"id\":123,\"name\":\"abcdefg\",\"params\":[\"hij\",\"klmn\",\"opqr\",\"stu\",\"vw\",\"xyz\"]}")

	r, _ = http.Get(ts.URL)
	err = gohttp.ResponseToError(r, nil, nil)
	assert.NoError(t, err)

	r, _ = http.Get(ts.URL)
	r.StatusCode = 404
	r.Status = "404 Not found"
	err = gohttp.ResponseToError(r, func(r *http.Response) bool { return r.StatusCode == 404 }, nil)
	assert.NoError(t, err)

	r, _ = http.Get(ts.URL)
	err = gohttp.ResponseToError(r,
		func(r *http.Response) bool { return r.StatusCode == 404 },
		func(i int, s1, s2 string) string {
			return fmt.Sprintf("%s %s %d", s2, s1, i)

		})
	assert.EqualError(t, err, "{\"id\":123,\"name\":\"abcdefg\",\"params\":[\"hij\",\"klmn\",\"opqr\",\"stu\",\"vw\",\"xyz\"]} 200 OK 200")
}
