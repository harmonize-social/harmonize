package main

import (
    "backend/internal/repositories"
    "backend/internal/routers"

    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt"
    "github.com/gorilla/mux"
)

func main() {
    repositories.CreateConnection()

    router := mux.NewRouter()
    //mount(router, "/api/v1", authedRoutes())
    mount(router, "/api/v1", unautherRoutes())

    fmt.Println("Starting server on port 8080")

    log.Fatal(http.ListenAndServe(":8080", router))
}

// Define some custom types were going to use within our tokens
type TokenContent struct {
    ID string
    expiry string
}

type CustomClaims struct {
    *jwt.StandardClaims
    TokenType string
    TokenContent
}

func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // TODO: Get auth token from header
        splitAuth := strings.Split(r.Header.Get("Authorization"), " ")
        if len(splitAuth) != 2 {
            fmt.Println("Invalid auth header")
            return
        }
        auth := splitAuth[1]
        var claims CustomClaims
        _, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte("secret"), nil
        })
        if err != nil {
            fmt.Println(err)
            return
        }
        r.Header.Add("id", claims.Id)
        next.ServeHTTP(w, r)
    })
}

func unautherRoutes() *mux.Router {
    router := mux.NewRouter()
    mount(router, "/session", routers.SessionRouter())
    return router
}

func authedRoutes() *mux.Router {
    router := mux.NewRouter()

    // oauth
    mount(router, "/oauth", routers.OAuthRouter())

    // user.go routers
    mount(router, "/user", routers.UserRouter())
    mount(router, "/follow", routers.FollowRouter())

    // interactions.go routers
    mount(router, "/post", routers.PostRouter())
    mount(router, "/like", routers.LikeRouter())
    mount(router, "/comment", routers.CommentRouter())
    mount(router, "/savedpost", routers.SavedPostRouter())
    router.Use(Middleware)
    return router
}

func mount(r *mux.Router, path string, handler http.Handler) {
    r.PathPrefix(path).Handler(
        http.StripPrefix(
            strings.TrimSuffix(path, "/"),
            handler,
        ),
    )
}
