package routers

import (
    "backend/internal/handlers"
    "backend/internal/models"
    "backend/internal/repositories"
    "fmt"
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt"
    "github.com/gorilla/mux"
)

/*
The full router that will be used by the server.
*/
func FullRouter() *mux.Router {
    router := mux.NewRouter()
    mount(router, "/api/v1/users", UserRouter())
    mount(router, "/api/v1", authedRoutes())
    return router
}

/*
Verify the JWT token and add the user id to the request header.
*/
func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        splitAuth := strings.Split(r.Header.Get("Authorization"), " ")
        if len(splitAuth) != 2 {
            models.Error(w, http.StatusUnauthorized, "Invalid auth header")
            // models.Error(w, http.StatusUnauthorized, err.Error())
            return
        }
        auth := splitAuth[1]
        token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
            return repositories.Secret, nil
        })
        if err != nil {
            models.Error(w, http.StatusUnauthorized, fmt.Sprint(token.Valid)+":"+err.Error())
            return
        }
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            r.Header.Add("id", fmt.Sprint(claims["id"]))
        } else {
            models.Error(w, http.StatusUnauthorized, "Invalid token")
        }
        next.ServeHTTP(w, r)
    })
}

/*
A router that requires authentication.
*/
func authedRoutes() *mux.Router {
    router := mux.NewRouter()

    mount(router, "/me", MeRouter())
    mount(router, "/oauth/callback", OAuthRouter())
    router.HandleFunc("/posts", handlers.GetUserPosts).Methods("GET")
    router.HandleFunc("/comment", handlers.CreateComment).Methods("POST")
    router.HandleFunc("/comment", handlers.DeleteComment).Methods("DELETE")
    router.HandleFunc("/comments", handlers.GetComments).Methods("GET")
    router.HandleFunc("/search", handlers.Search).Methods(http.MethodGet)

    router.Use(Middleware)
    return router
}

/*
A utility function to mount a router to a path.
*/
func mount(r *mux.Router, path string, handler http.Handler) {
    r.PathPrefix(path).Handler(
        http.StripPrefix(
            strings.TrimSuffix(path, "/"),
            handler,
        ),
    )
}
