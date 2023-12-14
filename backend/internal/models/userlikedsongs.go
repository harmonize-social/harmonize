package models

import (
    "github.com/google/uuid"
)

type UserLikedSongs struct {
	ID uuid.UUID `json:"id"`
    LibraryId uuid.UUID `json:"library_id"`
    SongId uuid.UUID `json:"song_id"`
}