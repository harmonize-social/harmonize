package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)

func LibraryRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/{service}/songs", handlers.SongsHandler).Methods("GET")
    router.HandleFunc("/{service}/artists", handlers.ArtistsHandler).Methods("GET")
    router.HandleFunc("/{service}/playlists", handlers.PlaylistHandler).Methods("GET")
    router.HandleFunc("/{service}/albums", handlers.AlbumHandler).Methods("GET")
    router.HandleFunc("/connected", handlers.ConnectedPlatforumsHandler).Methods("GET")
    router.HandleFunc("/unconnected", handlers.UnconnectedPlatformsHandler).Methods("GET")
    return router
}
