package models

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type ApiResponse struct {
    Error  string      `json:"error,omitempty"`
    Result interface{} `json:"result,omitempty"`
}

func Result(w http.ResponseWriter, result interface{}) {
    res := ApiResponse{
        Result: result,
    }

    err := json.NewEncoder(w).Encode(res)
    Error(w, http.StatusInternalServerError, fmt.Errorf("Error encoding JSON: %v", err).Error())
}

func Error(w http.ResponseWriter, code int, err string) {
    res := ApiResponse{
        Error: err,
    }

    w.WriteHeader(code)
    json.NewEncoder(w).Encode(res)
}
