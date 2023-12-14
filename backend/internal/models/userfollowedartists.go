package models

import (
    "github.com/google/uuid"
)

type UserFollowedArtists struct {
	ID uuid.UUID `json:"id"`
    LibraryId uuid.UUID `json:"library_id"`
    ArtistId uuid.UUID `json:"artist_id"`
}