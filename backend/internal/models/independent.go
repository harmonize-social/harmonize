package models

import "github.com/google/uuid"

/*
Represents the artist model which a post can have.

Not all fields are required, when this is not the root model.
*/
type Artist struct {
	ID uuid.UUID `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	MediaURL string `json:"mediaUrl,omitempty"`
}

/*
Represents the album model which a post can have.
*/
type Album struct {
	ID uuid.UUID `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Artists []Artist `json:"artists,omitempty"`
	Songs []Song `json:"songs,omitempty"`
	MediaURL string `json:"mediaUrl,omitempty"`
}

/*
Represents the playlist model which a post can have.
*/
type Playlist struct {
	ID uuid.UUID `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Songs []Song `json:"songs,omitempty"`
	MediaURL string `json:"mediaUrl,omitempty"`
}

/*
Represents the song model which a post can have.

Not all fields are required, when this is not the root model.
*/
type Song struct {
	ID uuid.UUID `json:"id,omitempty"`
	Album Album `json:"album,omitempty"`
	Title string `json:"title,omitempty"`
	PreviewURL string `json:"previewUrl,omitempty"`
	MediaURL string `json:"mediaUrl,omitempty"`
}
