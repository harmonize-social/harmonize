package models

type ApiResponse struct {
    Error string `json:"error:omitempty"`
    Value interface{} `json:"value:omitempty"`
}
