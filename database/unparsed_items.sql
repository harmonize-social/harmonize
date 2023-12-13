CREATE TABLE IF NOT EXISTS user_followed_artists(
    id uuid primary key,
    library_id uuid references libraries (id),
    artist_id uuid references platform_artists (id)
);

CREATE TABLE IF NOT EXISTS user_followed_albums(
    id uuid primary key,
    library_id uuid references libraries (id),
    album_id uuid references platform_albums (id)
);

CREATE TABLE IF NOT EXISTS user_followed_playlists(
    id uuid primary key,
    library_id uuid references libraries (id),
    playlist_id uuid references platform_playlists (id)
);

CREATE TABLE IF NOT EXISTS user_liked_songs(
    id uuid primary key,
    library_id uuid references libraries (id),
    song_id uuid references platform_songs (id)
);
