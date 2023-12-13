package routes

import (
    "net/http"

    "backend/internal/handlers" // Import your handlers package
)

// NewRouter initializes and returns a new mux router
func NewRouter() *http.ServeMux {
    mux := http.NewServeMux()

    // Define your routes here
    mux.HandleFunc("/oauth/spotify", handlers.OauthSpotify)
    mux.HandleFunc("/oauth/deezer", handlers.OauthDeezer)

    return mux
}
