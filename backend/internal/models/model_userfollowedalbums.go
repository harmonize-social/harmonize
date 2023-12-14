package models

import (
    "github.com/google/uuid"
)

type UserFollowedAlbums struct { // Is albums niet beter 'liked' ipv 'followed'?
	ID uuid.UUID `json:"id"`
    LibraryId uuid.UUID `json:"library_id"`
    AlbumId uuid.UUID `json:"album_id"`
}