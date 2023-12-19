package repositories

import "crypto/rand"

var Secret []byte

func GenerateSecret() (error) {
    // Define the length of the secret (in bytes)
    // Replace this value with the desired length of your secret
    secretLength := 32

    // Create a byte slice to store the secret
    Secret = make([]byte, secretLength)

    // Read random bytes to fill the secret
    _, err := rand.Read(Secret)
    return err
}
