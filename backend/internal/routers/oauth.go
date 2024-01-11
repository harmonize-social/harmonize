package routers

import (
    "backend/internal/platforms"

    "github.com/gorilla/mux"
)

/*
A router for the OAuth callbacks.
*/
func OAuthRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/deezer", platforms.DeezerCallback).Methods("GET")
    router.HandleFunc("/spotify", platforms.SpotifyCallback).Methods("GET")
    return router
}
