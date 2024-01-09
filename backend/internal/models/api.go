package models

import (
    "encoding/json"
    "fmt"
    "net/http"
)

/*
Represents a generic API response. If the request was successful, the Result
field will be populated with the result of the request. If the request failed,
the Error field will be populated with the error message.
*/
type ApiResponse struct {
    Error  string      `json:"error,omitempty"`
    Result interface{} `json:"result,omitempty"`
}

/*
Quick helper function to return a JSON response with the given result.
*/
func Result(w http.ResponseWriter, result interface{}) {
    res := ApiResponse{
        Result: result,
    }

    err := json.NewEncoder(w).Encode(res)
    if err == nil {
        return
    }
    Error(w, http.StatusInternalServerError, fmt.Errorf("Error encoding JSON: %v", err).Error())
}

/*
Quick helper function to return a JSON response with the given error message.
*/
func Error(w http.ResponseWriter, code int, err string) {
    res := ApiResponse{
        Error: err,
    }

    w.WriteHeader(code)
    json.NewEncoder(w).Encode(res)
}
