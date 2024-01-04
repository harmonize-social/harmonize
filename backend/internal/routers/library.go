package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)

func LibraryRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/{service}/{type}", handlers.LibraryHandler).Methods("GET")
    router.HandleFunc("/connected", handlers.ConnectedPlatforumsHandler).Methods("GET")
    router.HandleFunc("/unconnected", handlers.UnconnectedPlatformsHandler).Methods("GET")
    router.HandleFunc("/disconnect", handlers.DeleteConnectedPlatformHandler).Methods("DELETE")
    return router
}
