package models

import (
    "github.com/google/uuid"
)

type UserFollowedPlaylists struct {
	ID uuid.UUID `json:"id"`
    LibraryId uuid.UUID `json:"library_id"`
    PlaylistId uuid.UUID `json:"playlist_id"`
}