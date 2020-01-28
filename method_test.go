package gohttp_test

import (
	"testing"

	"github.com/marrbor/gohttp"
	"github.com/stretchr/testify/assert"
)

func TestHTTPMethod(t *testing.T) {
	var x gohttp.HTTPMethod
	x = gohttp.HttpMethods.GET
	assert.EqualValues(t, "GET", x.String())
	x = gohttp.HttpMethods.HEAD
	assert.EqualValues(t, "HEAD", x.String())
	x = gohttp.HttpMethods.POST
	assert.EqualValues(t, "POST", x.String())
	x = gohttp.HttpMethods.PUT
	assert.EqualValues(t, "PUT", x.String())
	x = gohttp.HttpMethods.PATCH
	assert.EqualValues(t, "PATCH", x.String())
	x = gohttp.HttpMethods.DELETE
	assert.EqualValues(t, "DELETE", x.String())
	x = gohttp.HttpMethods.CONNECT
	assert.EqualValues(t, "CONNECT", x.String())
	x = gohttp.HttpMethods.OPTIONS
	assert.EqualValues(t, "OPTIONS", x.String())
	x = gohttp.HttpMethods.TRACE
	assert.EqualValues(t, "TRACE", x.String())
}
