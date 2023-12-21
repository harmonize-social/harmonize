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
    mount(router, "/library", LibraryRouter())
    return router
}

func PostRouter() *mux.Router {
    router := mux.NewRouter()
    return router
}
