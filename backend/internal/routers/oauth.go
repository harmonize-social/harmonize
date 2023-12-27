package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)

func OAuthRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/deezer", handlers.DeezerCallback).Methods("GET")
    router.HandleFunc("/spotify", handlers.SpotifyCallback).Methods("GET")
    return router
}
