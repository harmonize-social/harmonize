package repositories

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/jackc/pgx/v4/pgxpool"
    "github.com/joho/godotenv"
)

// CreateConnection creates a connection pool with the database
func CreateConnection() (*pgxpool.Pool, error) {
    // Load .env file
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Create a new connection pool
    config, err := pgxpool.ParseConfig(os.Getenv("POSTGRES_URL"))
    if err != nil {
        return nil, fmt.Errorf("Unable to parse connection string: %v", err)
    }

    pool, err := pgxpool.ConnectConfig(context.Background(), config)
    if err != nil {
        return nil, fmt.Errorf("Unable to connect to database: %v", err)
    }

    // Check the connection
    if err := pool.Ping(context.Background()); err != nil {
        pool.Close()
        return nil, fmt.Errorf("Unable to ping database: %v", err)
    }

    fmt.Println("Connected to the database!")
    return pool, nil
}