package routers

import (
    "backend/internal/handlers"
    "github.com/gorilla/mux"
)

// Router is used in main.go
func UserRouter() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.CreateUser).Methods("POST")
    router.HandleFunc("/{id}", handlers.GetUser).Methods("GET")
    router.HandleFunc("/{id}", handlers.UpdateUser).Methods("PUT")
    router.HandleFunc("/{id}", handlers.DeleteUser).Methods("DELETE")
    return router
}

/* Just an example right now
func RouterAlbum(path string, r *mux.Router) {
    router := r

    router.HandleFunc(path+"/{id}", handlers.CreateAlbum).Methods("POST")
    router.HandleFunc(path+"/{id}", handlers.GetAlbum).Methods("GET", "OPTIONS")
    router.HandleFunc(path+"/{id}", handlers.UpdateAlbum).Methods("PUT", "OPTIONS")
    router.HandleFunc(path+"/{id}", handlers.DeleteAlbum).Methods("DELETE")
}*/
