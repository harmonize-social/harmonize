DROP TABLE IF EXISTS user_liked_songs;
DROP TABLE IF EXISTS user_followed_playlists;
DROP TABLE IF EXISTS user_liked_albums;
DROP TABLE IF EXISTS user_followed_artists;

DROP TABLE IF EXISTS saved_posts;
DROP TABLE IF EXISTS comment;
DROP TABLE IF EXISTS likes;
DROP TABLE IF EXISTS posts;

DROP TABLE IF EXISTS platform_playlists;
DROP TABLE IF EXISTS platform_songs;
DROP TABLE IF EXISTS platform_albums;
DROP TABLE IF EXISTS platform_artists;

DROP TABLE IF EXISTS playlist_songs;
DROP TABLE IF EXISTS playlists;
DROP TABLE IF EXISTS songs;
DROP TABLE IF EXISTS albums;
DROP TABLE IF EXISTS artists;

DROP TABLE IF EXISTS platforms;
DROP TABLE IF EXISTS images;
DROP TABLE IF EXISTS libraries;
DROP TABLE IF EXISTS connections;
DROP TABLE IF EXISTS follows;
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email VARCHAR(64) unique NOT NULL,
    username VARCHAR(64) unique NOT NULL,
    password_hash VARCHAR(256) NOT NULL
);


CREATE TABLE IF NOT EXISTS sessions (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
    expiry timestamptz NOT NULL
);


CREATE TABLE IF NOT EXISTS subscriptions (
    user_id UUID PRIMARY KEY REFERENCES users (id) NOT NULL,
    stripe_subscription VARCHAR(128) NOT NULL
);


CREATE TABLE IF NOT EXISTS follows(
    id UUID PRIMARY KEY,
    followed_id UUID REFERENCES users (id) NOT NULL,
    follower_id UUID REFERENCES users (id) NOT NULL,
    date timestamptz NOT NULL
);


CREATE TABLE IF NOT EXISTS connections(
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
    access_token VARCHAR(1024) NOT NULL,
    refresh_token VARCHAR(512) NOT NULL,
    expiry timestamptz NOT NULL
);


CREATE TABLE IF NOT EXISTS libraries(
    id UUID PRIMARY KEY,
    platform_id VARCHAR(64) NOT NULL,
    connection_id UUID REFERENCES connections (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS images(
    id UUID PRIMARY KEY,
    path VARCHAR(64) NOT NULL
);


CREATE TABLE IF NOT EXISTS platforms(
    id UUID PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    icon_id UUID REFERENCES images (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS artists(
    id UUID PRIMARY KEY,
    name VARCHAR(128) NOT NULL
);


CREATE TABLE IF NOT EXISTS albums(
    id UUID PRIMARY KEY,
    author UUID NOT NULL REFERENCES artists (id),
    title VARCHAR(128) NOT NULL
);


CREATE TABLE IF NOT EXISTS playlists(
    id UUID PRIMARY KEY,
    title VARCHAR(128) NOT NULL
);


CREATE TABLE IF NOT EXISTS songs(
    id UUID PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    album_id UUID REFERENCES albums (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS playlist_songs(
    id UUID PRIMARY KEY,
    playlist_id UUID REFERENCES playlists (id) NOT NULL,
    song_id UUID REFERENCES songs (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS platform_artists(
    id UUID PRIMARY KEY,
    platform_specific_id VARCHAR(64) NOT NULL,
    artist_id UUID REFERENCES artists (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS platform_albums(
    id UUID PRIMARY KEY,
    platform_specific_id VARCHAR(64) NOT NULL,
    album_id UUID REFERENCES albums (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS platform_songs(
    id UUID PRIMARY KEY,
    platform_specific_id VARCHAR(64) NOT NULL,
    song_id UUID REFERENCES songs (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS platform_playlists(
    id UUID PRIMARY KEY,
    platform_specific_id VARCHAR(64) NOT NULL,
    playlist_id UUID REFERENCES playlists (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS posts(
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
    caption VARCHAR(256) NOT NULL,
    type VARCHAR(16) NOT NULL,
    type_specific_id UUID
);


CREATE TABLE IF NOT EXISTS likes(
    id UUID PRIMARY KEY,
    post_id UUID REFERENCES posts (id) NOT NULL,
    user_id UUID REFERENCES users (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS comment(
    id UUID PRIMARY KEY,
    post_id UUID REFERENCES posts (id) NOT NULL,
    user_id UUID REFERENCES users (id) NOT NULL,
    reply_to_id UUID REFERENCES comment (id) NOT NULL,
    message VARCHAR(256) NOT NULL
);


CREATE TABLE IF NOT EXISTS saved_posts(
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES posts (id) NOT NULL,
    post_id UUID REFERENCES posts (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS user_followed_artists(
    id UUID PRIMARY KEY,
    library_id UUID REFERENCES libraries (id) NOT NULL,
    artist_id UUID REFERENCES platform_artists (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS user_liked_albums(
    id UUID PRIMARY KEY,
    library_id UUID REFERENCES libraries (id) NOT NULL,
    album_id UUID REFERENCES platform_albums (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS user_followed_playlists(
    id UUID PRIMARY KEY,
    library_id UUID REFERENCES libraries (id) NOT NULL,
    playlist_id UUID REFERENCES platform_playlists (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS user_liked_songs(
    id UUID PRIMARY KEY,
    library_id UUID REFERENCES libraries (id) NOT NULL,
    song_id UUID REFERENCES platform_songs (id) NOT NULL
);
