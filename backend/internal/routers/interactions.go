package routers

import (
    "backend/internal/handlers"
    
    "github.com/gorilla/mux"
)

func PostRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.CreatePost).Methods("POST")
    router.HandleFunc("/{id}", handlers.GetPost).Methods("GET")
    router.HandleFunc("/{id}", handlers.UpdatePost).Methods("PUT")
    router.HandleFunc("/{id}", handlers.DeletePost).Methods("DELETE")
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