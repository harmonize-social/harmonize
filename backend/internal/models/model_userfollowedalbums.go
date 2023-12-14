package models

import (
    "github.com/google/uuid"
)

type UserLikedAlbums struct {
	ID uuid.UUID `json:"id"`
    LibraryId uuid.UUID `json:"library_id"`
    AlbumId uuid.UUID `json:"album_id"`
}