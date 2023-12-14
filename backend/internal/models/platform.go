package models

import (
    "github.com/google/uuid"
)

type User struct {
    ID uuid.UUID `json:"id"`
    Email string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password_hash"`
}

type UserLikedAlbums struct {
    ID uuid.UUID `json:"id"`
    LibraryId uuid.UUID `json:"library_id"`
    AlbumId uuid.UUID `json:"album_id"`
}

type UserFollowedArtists struct {
    ID uuid.UUID `json:"id"`
    LibraryId uuid.UUID `json:"library_id"`
    ArtistId uuid.UUID `json:"artist_id"`
}

type UserFollowedPlaylists struct {
    ID uuid.UUID `json:"id"`
    LibraryId uuid.UUID `json:"library_id"`
    PlaylistId uuid.UUID `json:"playlist_id"`
}

type UserLikedSongs struct {
    ID uuid.UUID `json:"id"`
    LibraryId uuid.UUID `json:"library_id"`
    SongId uuid.UUID `json:"song_id"`
}
