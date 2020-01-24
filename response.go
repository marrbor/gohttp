package gohttp

import (
	"encoding/json"
	"net/http"
)

////// Send response (for server side program)

// ResponseOK returns 200 ok.
func ResponseOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

// errResponse returns error response.
func errResponse(w http.ResponseWriter, code int, err error) {
	msg := http.StatusText(code)
	if err != nil {
		m := err.Error()
		if m != "" {
			msg = m
		}
	}
	http.Error(w, msg, code)
}

// BadRequest returns http 400
func BadRequest(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusBadRequest, err)
}

// Unauthorized returns http 401
func Unauthorized(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusUnauthorized, err)
}

// Forbidden returns http 403
func Forbidden(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusForbidden, err)
}

// NotFound returns http 404
func NotFound(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusNotFound, err)
}

// MethodNotAllowed returns http 405
func MethodNotAllowed(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusMethodNotAllowed, err)
}

// InternalServerError returns http 500
func InternalServerError(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusInternalServerError, err)
}

// JSONResponse returns JSON object
func JSONResponse(w http.ResponseWriter, data interface{}) error {
	if data == nil {
		ResponseOK(w)
		return nil
	}

	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(j); err != nil {
		return err
	}
	return nil
}
