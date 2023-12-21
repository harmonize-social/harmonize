package repositories

import (
    "crypto/rand"
    "encoding/hex"
    "os"
)

var Secret []byte

func GenerateSecret() (error) {
    EnvSecret := os.Getenv("BACKEND_SECRET")
    if EnvSecret != "" {
        Secret, _ = hex.DecodeString(EnvSecret)
        return nil
    }
    // Define the length of the secret (in bytes)
    // Replace this value with the desired length of your secret
    secretLength := 32

    // Create a byte slice to store the secret
    Secret = make([]byte, secretLength)

    // Read random bytes to fill the secret
    _, err := rand.Read(Secret)
    return err
}
