package main

import (
    "backend/internal/routers"

    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"

    "github.com/gorilla/mux"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func main() {
    // Load .env file
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

    Pool = dbpool

	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

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
