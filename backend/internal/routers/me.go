package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)


func MeRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/feed", handlers.GetFeed).Methods("GET")
    router.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
    router.HandleFunc("/saved", handlers.GetSavedPosts).Methods("GET")
    router.HandleFunc("/saved", handlers.PostSavedPost).Methods("POST")
    router.HandleFunc("/saved", handlers.DeleteSavedPost).Methods("DELETE")
    router.HandleFunc("/liked", handlers.GetLikedPosts).Methods("GET")
    router.HandleFunc("/liked", handlers.PostLikedPost).Methods("POST")
    router.HandleFunc("/liked", handlers.DeleteLikedPost).Methods("DELETE")
    mount(router, "/library", LibraryRouter())
    return router
}
