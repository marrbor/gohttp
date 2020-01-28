// wrap net/http.Method
package gohttp

import "net/http"

type HTTPMethod struct{ method string }

func (h HTTPMethod) String() string {
	return h.method
}

var HttpMethods = struct {
	GET     HTTPMethod
	HEAD    HTTPMethod
	POST    HTTPMethod
	PUT     HTTPMethod
	PATCH   HTTPMethod
	DELETE  HTTPMethod
	CONNECT HTTPMethod
	OPTIONS HTTPMethod
	TRACE   HTTPMethod
}{
	GET:     HTTPMethod{http.MethodGet},
	HEAD:    HTTPMethod{http.MethodHead},
	POST:    HTTPMethod{http.MethodPost},
	PUT:     HTTPMethod{http.MethodPut},
	PATCH:   HTTPMethod{http.MethodPatch},
	DELETE:  HTTPMethod{http.MethodDelete},
	CONNECT: HTTPMethod{http.MethodConnect},
	OPTIONS: HTTPMethod{http.MethodOptions},
	TRACE:   HTTPMethod{http.MethodTrace},
}
