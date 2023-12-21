package handlers

import (
    _ "github.com/lib/pq"      // postgres golang driver
    "net/http" // used to access the request and response object of the api
)

type response struct {
    ID      string `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}

// setCommonHeaders sets common headers for CORS
func setCommonHeaders(w http.ResponseWriter) {
}

// setAdditionalHeaders sets additional headers specific to some handler functions
func setAdditionalHeaders(w http.ResponseWriter, methods ...string) {
}
