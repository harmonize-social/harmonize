package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)

func OAuthRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/callback/deezer", handlers.DeezerCallback).Methods("GET")
    return router
}
