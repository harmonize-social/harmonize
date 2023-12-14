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

    /* Just an example right now
       // Album routes
       path := "/api/album"
       routers.RouterAlbum(path, router)
    */

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
