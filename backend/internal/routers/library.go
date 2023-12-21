package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)

func LibraryRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/{service}/songs", handlers.SongsHandler).Methods("GET")
    router.HandleFunc("/{service}/artists", handlers.ArtistsHandler).Methods("GET")
    return router
}
