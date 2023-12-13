CREATE TABLE IF NOT EXISTS platform_artists(
    id uuid primary key,
    platform_specific_id varchar(64),
    artist_id uuid references artists (id)
);

CREATE TABLE IF NOT EXISTS platform_albums(
    id uuid primary key,
    platform_specific_id varchar(64),
    album_id uuid references albums (id)
);

CREATE TABLE IF NOT EXISTS platform_songs(
    id uuid primary key,
    platform_specific_id varchar(64),
    song_id uuid references songs (id)
);

CREATE TABLE IF NOT EXISTS platform_playlists(
    id uuid primary key,
    platform_specific_id varchar(64),
    playlist_id uuid references playlists (id)
);
