package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)

func MeRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/feed", handlers.GetFeed).Methods("GET")
    router.HandleFunc("/posts", handlers.GetMePosts).Methods("GET")
    router.HandleFunc("/posts", handlers.NewPost).Methods("POST")
    router.HandleFunc("/saved", handlers.GetSavedPosts).Methods("GET")
    router.HandleFunc("/saved", handlers.PostSavedPost).Methods("POST")
    router.HandleFunc("/saved", handlers.DeleteSavedPost).Methods("DELETE")
    router.HandleFunc("/liked", handlers.GetLikedPosts).Methods("GET")
    router.HandleFunc("/liked", handlers.PostLikedPost).Methods("POST")
    router.HandleFunc("/liked", handlers.DeleteLikedPost).Methods("DELETE")
    router.HandleFunc("/follow", handlers.PostFollow).Methods("POST")
    router.HandleFunc("/follow", handlers.DeleteFollow).Methods("DELETE")
    router.HandleFunc("/following", handlers.GetFollowing).Methods("GET")
    router.HandleFunc("/followers", handlers.GetFollowers).Methods("GET")
    mount(router, "/library", LibraryRouter())
    return router
}
