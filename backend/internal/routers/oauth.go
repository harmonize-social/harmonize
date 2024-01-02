package routers

import (
    "backend/internal/handlers"
    "backend/internal/platforms"

    "github.com/gorilla/mux"
)

func OAuthRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/deezer", handlers.DeezerCallback).Methods("GET")
    router.HandleFunc("/spotify", platforms.SpotifyCallback).Methods("GET")
    return router
}
