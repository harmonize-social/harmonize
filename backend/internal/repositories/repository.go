package repositories

import (
	"github.com/joho/godotenv"
	"os" // used to read the environment variable
	_ "github.com/lib/pq"      // postgres golang driver
	"fmt"
	"log"
	"database/sql"
)

// Create connection with db
func CreateConnection() *sql.DB {
    // load .env file
    err := godotenv.Load(".env")

    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Open the connection
    db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

    if err != nil {
        panic(err)
    }

    // check the connection
    err = db.Ping()

    if err != nil {
        panic(err)
    }

    fmt.Println("Connected to database!")
    // return the connection
    return db
}