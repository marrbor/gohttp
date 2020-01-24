package gohttp

import (
	"net/http"
	"strings"
)

// PreflightForCORS sets HTTP Headers needed for Preflight For CORS to given ResponseWriter.
// PreflightForCORS returns true when add header, false when nothing to do.
// When nil specified for allowHeaders and/or allowMethods, PreflightForCORS sets them as same as client request.
// refer:
// https://medium.com/google-cloud-jp/https-medium-com-google-cloud-jp-gae-api-12e5ab274719
// https://symfoware.blog.fc2.com/blog-entry-2010.html
// https://qiita.com/rooooomania/items/4d0f6275372f413765de
func PreflightForCORS(w http.ResponseWriter, r *http.Request, allowHeaders, allowMethods *[]string, allowOrigin string) bool {
	if r.Method != http.MethodOptions {
		return false
	}

	// When nil specified for allowHeaders, PreflightForCORS sets them as same as client request.
	headers := r.Header.Get("Access-Control-Request-Headers")
	if allowHeaders != nil {
		headers = strings.Join(*allowHeaders, ", ")
	}

	// Also when nil specified for allowMethods, PreflightForCORS sets them as same as client request.
	methods := r.Header.Get("Access-Control-Request-Method")
	if allowMethods != nil {
		methods = strings.Join(*allowMethods, ", ")
	}

	w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
	w.Header().Set("Access-Control-Allow-Headers", headers)
	w.Header().Set("Access-Control-Allow-Methods", methods)
	ResponseOK(w)
	return true
}

// AddCORSHeader add a header for authentication.
func AddCORSHeader(w http.ResponseWriter, allowOrigin string) {
	w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
}
