package routers

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt"
    "github.com/gorilla/mux"
)

func FullRouter() *mux.Router {
    router := mux.NewRouter()
    mount(router, "/api/v1/session", SessionRouter())
    mount(router, "/api/v1", authedRoutes())
    return router
}

func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        splitAuth := strings.Split(r.Header.Get("Authorization"), " ")
        if len(splitAuth) != 2 {
            fmt.Println("Invalid auth header")
            return
        }
        auth := splitAuth[1]
        var claims jwt.MapClaims
        _, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte("secret"), nil
        })
        if err != nil {
            fmt.Println(err)
            return
        }
        r.Header.Add("id", claims["id"].(string))
        next.ServeHTTP(w, r)
    })
}

func authedRoutes() *mux.Router {
    router := mux.NewRouter()

    // oauth
    mount(router, "/oauth", OAuthRouter())

    // user.go routers
    mount(router, "/user", UserRouter())
    mount(router, "/follow", FollowRouter())

    // interactions.go routers
    mount(router, "/post", PostRouter())
    mount(router, "/like", LikeRouter())
    mount(router, "/comment", CommentRouter())
    mount(router, "/savedpost", SavedPostRouter())
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
