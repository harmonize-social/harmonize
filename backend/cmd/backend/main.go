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

    mount(router, "/api/session", routers.SessionRouter())

    // oauth
    mount(router, "/api/oauth", routers.OAuthRouter())

    // user.go routers
    mount(router, "/api/user", routers.UserRouter())
    mount(router, "/api/follow", routers.FollowRouter())

    // interactions.go routers
    mount(router, "/api/post", routers.PostRouter())
    mount(router, "/api/like", routers.LikeRouter())
    mount(router, "/api/comment", routers.CommentRouter())
    mount(router, "/api/savedpost", routers.SavedPostRouter())

    fmt.Println("Starting server on port 8080")

    log.Fatal(http.ListenAndServe(":8080", router))
}

func mount(r *mux.Router, path string, handler http.Handler) {
    r.PathPrefix(path).Handler(
        http.StripPrefix(
            strings.TrimSuffix(path, "/"),
            handler,
        ),
    )
}
