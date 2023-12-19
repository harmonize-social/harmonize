package routers

import (
    "backend/internal/handlers"

    "github.com/gorilla/mux"
)

func UserRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.CreateUser).Methods("POST")
    router.HandleFunc("/{id}", handlers.GetUser).Methods("GET")
    router.HandleFunc("/{id}", handlers.UpdateUser).Methods("PUT")
    router.HandleFunc("/{id}", handlers.DeleteUser).Methods("DELETE")
    return router
}

func SessionRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.CreateSession).Methods("POST")
    router.HandleFunc("/{id}", handlers.GetSession).Methods("GET")
    router.HandleFunc("/{id}", handlers.UpdateSession).Methods("PUT")
    router.HandleFunc("/{id}", handlers.DeleteSession).Methods("DELETE")
    return router
}

func FollowRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.CreateFollow).Methods("POST")
    router.HandleFunc("/{id}", handlers.GetFollow).Methods("GET")
    router.HandleFunc("/{id}", handlers.DeleteFollow).Methods("DELETE")
    return router
}