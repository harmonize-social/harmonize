package routers

import (
	"github.com/gorilla/mux"
	"backend/internal/handlers"
)

// Router is used in main.go
func RouterUser(path string, r *mux.Router) {
	router := r

	router.HandleFunc(path+"/{id}", handlers.CreateUser).Methods("POST")
	router.HandleFunc(path+"/{id}", handlers.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc(path+"/{id}", handlers.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc(path+"/{id}", handlers.DeleteUser).Methods("DELETE")
}

/* Just an example right now
func RouterAlbum(path string, r *mux.Router) {
	router := r

	router.HandleFunc(path+"/{id}", handlers.CreateAlbum).Methods("POST")
	router.HandleFunc(path+"/{id}", handlers.GetAlbum).Methods("GET", "OPTIONS")
	router.HandleFunc(path+"/{id}", handlers.UpdateAlbum).Methods("PUT", "OPTIONS")
	router.HandleFunc(path+"/{id}", handlers.DeleteAlbum).Methods("DELETE")
}*/