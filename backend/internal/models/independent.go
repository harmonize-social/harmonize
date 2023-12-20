package models

import "github.com/google/uuid"

type Artist struct {
    ID uuid.UUID `json:"id"`
    Name string `json:"name"`
    MediaURL string `json:"mediaUrl"`
}

type Album struct {
    ID uuid.UUID `json:"id"`
    Title string `json:"title"`
    Artists []Artist `json:"artists"`
    Songs []Song `json:"songs"`
    MediaURL string `json:"mediaUrl"`
}

type Playlist struct {
    ID uuid.UUID `json:"id"`
    Title string `json:"title"`
    Songs []Song `json:"songs"`
    MediaURL string `json:"mediaUrl"`
}

type Song struct {
    ID uuid.UUID `json:"id"`
    Artists []Artist `json:"artists"`
    Title string `json:"title"`
    MediaURL string `json:"mediaUrl"`
}
