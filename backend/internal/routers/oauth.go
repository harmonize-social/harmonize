package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)

func OAuthRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/spotify", handlers.OauthSpotify).Methods("GET")
    router.HandleFunc("/deezer", handlers.OauthDeezer).Methods("GET")
    router.HandleFunc("/callback", handlers.OauthCallback).Methods("GET")
    return router
}
