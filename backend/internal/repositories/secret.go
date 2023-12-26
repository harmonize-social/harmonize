package repositories

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "os"
)

var Secret []byte

var SpotifySecret string
var SpotifyClientId string
var SpotifyRedirect string

var DeezerSecret string
var DeezerClientId string
var DeezerRedirect string

func LoadSecret() error {
    EnvSecret := os.Getenv("BACKEND_SECRET")
    if EnvSecret == "" {
        return nil
    }
    Secret, _ = hex.DecodeString(EnvSecret)
    return nil
}

func GenerateSecret() error {
    secretLength := 32
    Secret = make([]byte, secretLength)
    _, err := rand.Read(Secret)
    return err
}

func LoadSpotifyEnv() error {
    SpotifySecret = os.Getenv("SPOTIFY_SECRET")
    SpotifyClientId = os.Getenv("SPOTIFY_CLIENT_ID")
    if SpotifySecret == "" || SpotifyClientId == "" {
        return fmt.Errorf("SPOTIFY_SECRET or SPOTIFY_CLIENT_ID not set")
    }
    return nil
}

func LoadDeezerEnv() error {
    DeezerSecret = os.Getenv("DEEZER_SECRET")
    DeezerClientId = os.Getenv("DEEZER_CLIENT_ID")
    if DeezerSecret == "" || DeezerClientId == "" {
        return fmt.Errorf("DEEZER_SECRET or DEEZER_CLIENT_ID not set")
    }
    return nil
}

func LoadCallbackURLs() error {
    SpotifyRedirect = os.Getenv("SPOTIFY_REDIRECT")
    DeezerRedirect = os.Getenv("DEEZER_REDIRECT")
    if SpotifyRedirect == "" || DeezerRedirect == "" {
        return fmt.Errorf("SPOTIFY_REDIRECT or DEEZER_REDIRECT not set")
    }
    return nil
}

func LoadEnv() {
    err := LoadSecret()
    if err != nil {
        fmt.Println("BACKEND_SECRET not set, generating new secret")
        err = GenerateSecret()
        if err != nil {
            fmt.Println("Failed to generate secret")
            os.Exit(1)
        }
    }
    err = LoadSpotifyEnv()
    if err != nil {
        fmt.Println("Warning: Spotify environment variables not set")
    }
    err = LoadDeezerEnv()
    if err != nil {
        fmt.Println("Warning: Deezer environment variables not set")
    }
    err = LoadCallbackURLs()
    if err != nil {
        fmt.Println("Warning: Callback URLs not set")
    }
}
