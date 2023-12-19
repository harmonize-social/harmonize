package main

import (
    "backend/internal/repositories"
    "backend/internal/routers"

    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/gorilla/mux"
)

func main() {
    repositories.CreateConnection()

    router := mux.NewRouter()
    mount(router, "/api/v1/", unautherRoutes())
    mount(router, "/api/v1/", authedRoutes())

    fmt.Println("Starting server on port 8080")

    log.Fatal(http.ListenAndServe(":8080", router))
}

func unautherRoutes() *mux.Router {
    router := mux.NewRouter()
    mount(router, "/session", routers.SessionRouter())
    return router
}

func authedRoutes() *mux.Router {
    router := mux.NewRouter()

    mount(router, "/session", routers.SessionRouter())

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
