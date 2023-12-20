package repositories

import (
    "context"
    "fmt"
    "os"

    "github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool

// CreateConnection creates a connection pool with the database
func CreateConnection() error {
    // Create a new connection pool
    config, err := pgxpool.ParseConfig(os.Getenv("POSTGRES_URL"))
    if err != nil {
        return fmt.Errorf("Unable to parse connection string: %v", err)
    }

    pool, err := pgxpool.ConnectConfig(context.Background(), config)
    if err != nil {
        return fmt.Errorf("Unable to connect to database: %v", err)
    }

    // Check the connection
    if err := pool.Ping(context.Background()); err != nil {
        pool.Close()
        return fmt.Errorf("Unable to ping database: %v", err)
    }

    Pool = pool

    fmt.Println("Connected to the database!")
    return nil
}
