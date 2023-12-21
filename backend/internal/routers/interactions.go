package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)


func MeRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/feed", handlers.GetFeed).Methods("GET")
    router.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
    mount(router, "/library", LibraryRouter())
    return router
}

func PostRouter() *mux.Router {
    router := mux.NewRouter()
    return router
}

func LikeRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.CreateLike).Methods("POST")
    router.HandleFunc("/{id}", handlers.GetLike).Methods("GET")
    router.HandleFunc("/{id}", handlers.DeleteLike).Methods("DELETE")
    return router
}

func CommentRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.CreateComment).Methods("POST")
    router.HandleFunc("/{id}", handlers.GetComment).Methods("GET")
    router.HandleFunc("/{id}", handlers.UpdateComment).Methods("PUT")
    router.HandleFunc("/{id}", handlers.DeleteComment).Methods("DELETE")
    return router
}

func SavedPostRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.CreateSavedPost).Methods("POST")
    router.HandleFunc("/{id}", handlers.GetSavedPost).Methods("GET")
    router.HandleFunc("/{id}", handlers.UpdateSavedPost).Methods("PUT")
    router.HandleFunc("/{id}", handlers.DeleteSavedPost).Methods("DELETE")
    return router
}
