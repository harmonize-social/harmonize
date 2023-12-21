package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)

func OAuthRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/spotify", handlers.SpotifyURL).Methods("GET")
    router.HandleFunc("/deezer", handlers.DeezerURL).Methods("GET")
    router.HandleFunc("/callback/deezer", handlers.DeezerCallback).Methods("GET")
    return router
}
