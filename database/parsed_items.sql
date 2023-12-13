CREATE TABLE IF NOT EXISTS artists(
    id uuid primary key,
    name varchar(128) not null
);

CREATE TABLE IF NOT EXISTS albums(
    id uuid primary key,
    author uuid not null references artists (id),
    title varchar(128) not null
);

CREATE TABLE IF NOT EXISTS playlists(
    id uuid primary key,
    title varchar(128) not null
);

CREATE TABLE IF NOT EXISTS playlist_songs(
    id uuid primary key,
    playlist_id uuid references playlists (id),
    song_id uuid references songs (id),
);

CREATE TABLE IF NOT EXISTS songs(
    id uuid primary key,
    title varchar(128),
    album_id uuid references albums (id),
);
