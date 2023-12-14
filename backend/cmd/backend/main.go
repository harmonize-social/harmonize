package main

import (
    "backend/internal/routers"
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    // User routes
    //
    mount(router, "/api/user", routers.UserRouter())
    mount(router, "/api/oauth", routers.OAuthRouter())
    mount(router, "/api/post", routers.PostRouter())
    mount(router, "/api/like", routers.LikeRouter())

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
