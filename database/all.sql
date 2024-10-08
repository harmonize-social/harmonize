DROP EXTENSION IF EXISTS "pg_trgm";
DROP EXTENSION IF EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS user_liked_songs;
DROP TABLE IF EXISTS user_followed_playlists;
DROP TABLE IF EXISTS user_liked_albums;
DROP TABLE IF EXISTS user_followed_artists;

DROP TABLE IF EXISTS saved_posts;
DROP TABLE IF EXISTS liked_posts;
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS likes;
DROP TABLE IF EXISTS posts;

DROP TABLE IF EXISTS platform_playlists;
DROP TABLE IF EXISTS platform_songs;
DROP TABLE IF EXISTS platform_albums;
DROP TABLE IF EXISTS platform_artists;

DROP TABLE IF EXISTS playlist_songs;
DROP TABLE IF EXISTS playlists;
DROP TABLE IF EXISTS songs;
DROP TABLE IF EXISTS artists_album;
DROP TABLE IF EXISTS artists;
DROP TABLE IF EXISTS albums;

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
    email VARCHAR(1024) unique NOT NULL,
    username VARCHAR(1024) unique NOT NULL,
    password_hash VARCHAR(1024) NOT NULL
);


CREATE TABLE IF NOT EXISTS sessions (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
    expiry timestamptz NOT NULL
);


CREATE TABLE IF NOT EXISTS subscriptions (
    user_id UUID PRIMARY KEY REFERENCES users (id) NOT NULL,
    stripe_subscription VARCHAR(1024) NOT NULL
);


CREATE TABLE IF NOT EXISTS follows(
    id UUID PRIMARY KEY,
    followed_id UUID REFERENCES users (id) NOT NULL,
    follower_id UUID REFERENCES users (id) NOT NULL,
    CONSTRAINT follows_unique_user_combo UNIQUE (followed_id, follower_id),
    date timestamptz NOT NULL
);


CREATE TABLE IF NOT EXISTS connections(
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
    access_token VARCHAR(1024) NOT NULL,
    refresh_token VARCHAR(1024) NOT NULL,
    expiry timestamptz NOT NULL
);


CREATE TABLE IF NOT EXISTS libraries(
    id UUID PRIMARY KEY,
    platform_id VARCHAR(1024) NOT NULL,
    connection_id UUID REFERENCES connections (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS images(
    id UUID PRIMARY KEY,
    path VARCHAR(1024) NOT NULL
);


CREATE TABLE IF NOT EXISTS platforms(
    id VARCHAR(1024) PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    icon_id UUID REFERENCES images (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS artists(
    id UUID PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    media_url VARCHAR(1024) NOT NULL
);

CREATE TABLE IF NOT EXISTS albums(
    id UUID PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    media_url VARCHAR(1024) NOT NULL
);

CREATE TABLE IF NOT EXISTS artists_album(
    id UUID PRIMARY KEY,
    artist_id UUID REFERENCES artists (id) NOT NULL,
    album_id UUID REFERENCES albums (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS songs(
    id UUID PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    album_id UUID REFERENCES albums (id) NOT NULL,
    media_url VARCHAR(1024) NOT NULL,
    preview_url VARCHAR(1024) NOT NULL
);

CREATE TABLE IF NOT EXISTS playlists(
    id UUID PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    media_url VARCHAR(1024) NOT NULL
);


CREATE TABLE IF NOT EXISTS playlist_songs(
    id UUID PRIMARY KEY,
    playlist_id UUID REFERENCES playlists (id) NOT NULL,
    song_id UUID REFERENCES songs (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS platform_artists(
    id UUID PRIMARY KEY,
    platform_specific_id VARCHAR(1024) NOT NULL,
    platform_id VARCHAR(1024) REFERENCES platforms(id) NOT NULL,
    artist_id UUID REFERENCES artists (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS platform_albums(
    id UUID PRIMARY KEY,
    platform_specific_id VARCHAR(1024) NOT NULL,
    platform_id VARCHAR(1024) REFERENCES platforms(id) NOT NULL,
    album_id UUID REFERENCES albums (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS platform_songs(
    id UUID PRIMARY KEY,
    platform_specific_id VARCHAR(1024) NOT NULL,
    platform_id VARCHAR(1024) REFERENCES platforms(id) NOT NULL,
    song_id UUID REFERENCES songs (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS platform_playlists(
    id UUID PRIMARY KEY,
    platform_specific_id VARCHAR(1024) NOT NULL,
    platform_id VARCHAR(1024) REFERENCES platforms(id) NOT NULL,
    playlist_id UUID REFERENCES playlists (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS posts(
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
    created_at timestamptz NOT NULL,
    caption VARCHAR(1024) NOT NULL,
    type VARCHAR(1024) NOT NULL,
    type_specific_id UUID
);


CREATE TABLE IF NOT EXISTS likes(
    id UUID PRIMARY KEY,
    post_id UUID REFERENCES posts (id) NOT NULL,
    user_id UUID REFERENCES users (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS comments(
    id UUID PRIMARY KEY,
    post_id UUID REFERENCES posts (id) NOT NULL,
    username varchar(1024) REFERENCES users (username) NOT NULL,
    reply_to_id UUID,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    message VARCHAR(1024) NOT NULL
);

CREATE TABLE IF NOT EXISTS saved_posts (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
    post_id UUID REFERENCES posts (id) NOT NULL,
    CONSTRAINT saved_unique_user_post_combo UNIQUE (user_id, post_id)
);

CREATE TABLE IF NOT EXISTS liked_posts (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
    post_id UUID REFERENCES posts (id) NOT NULL,
    CONSTRAINT liked_unique_user_post_combo UNIQUE (user_id, post_id)
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


DROP FUNCTION IF EXISTS insert_new_artist;
CREATE OR REPLACE FUNCTION insert_new_artist(new_platform_id VARCHAR(1024), platform_specific_id_input VARCHAR(1024), new_name VARCHAR(1024), new_media_url VARCHAR(1024))
RETURNS TABLE (value1 UUID, value2 UUID) AS $$
DECLARE
    artists_artist_id UUID;
    platform_artists_artist_id UUID;
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM platform_artists WHERE platform_specific_id = platform_specific_id_input
    ) THEN
        INSERT INTO artists (id, name, media_url)
        VALUES (uuid_generate_v4(), new_name, new_media_url)
        RETURNING id INTO artists_artist_id;

        INSERT INTO platform_artists (id, platform_id, platform_specific_id, artist_id)
        VALUES (uuid_generate_v4(), new_platform_id, platform_specific_id_input, artists_artist_id)
        RETURNING id INTO platform_artists_artist_id;

        RETURN QUERY SELECT artists_artist_id, platform_artists_artist_id;
    ELSE
        SELECT artists.id FROM platform_artists
        JOIN artists ON platform_artists.artist_id = artists.id
        WHERE platform_specific_id = platform_specific_id_input INTO artists_artist_id;

        SELECT platform_artists.id FROM platform_artists
        WHERE platform_specific_id = platform_specific_id_input INTO platform_artists_artist_id;

        RETURN QUERY SELECT artists_artist_id, platform_artists_artist_id;
    END IF;
END;
$$ LANGUAGE plpgsql;


DROP FUNCTION IF EXISTS insert_new_album;
CREATE OR REPLACE FUNCTION insert_new_album(new_platform_id VARCHAR(1024), platform_specific_album_id VARCHAR(1024), new_name VARCHAR(1024), new_media_url VARCHAR(1024))
RETURNS TABLE (value1 UUID, value2 UUID) AS $$
DECLARE
    albums_album_id UUID;
    platform_albums_album_id UUID;
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM platform_albums WHERE platform_specific_id = platform_specific_album_id
    ) THEN
        INSERT INTO albums (id, name, media_url)
        VALUES (uuid_generate_v4(), new_name, new_media_url)
        RETURNING id INTO albums_album_id;

        INSERT INTO platform_albums (id, platform_id, platform_specific_id, album_id)
        VALUES (uuid_generate_v4(), new_platform_id, platform_specific_album_id, albums_album_id)
        RETURNING id INTO platform_albums_album_id;

        RETURN QUERY SELECT albums_album_id, platform_albums_album_id;
    ELSE
        SELECT albums.id FROM platform_albums
        JOIN albums ON platform_albums.album_id = albums.id
        WHERE platform_specific_id = platform_specific_album_id INTO albums_album_id;

        SELECT platform_albums.id FROM platform_albums
        WHERE platform_specific_id = platform_specific_album_id INTO platform_albums_album_id;

        RETURN QUERY SELECT albums_album_id, platform_albums_album_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

DROP FUNCTION IF EXISTS insert_new_playlist;
CREATE OR REPLACE FUNCTION insert_new_playlist(new_platform_id VARCHAR(1024), platform_specific_id_input VARCHAR(1024), new_name VARCHAR(1024), new_media_url VARCHAR(1024))
RETURNS TABLE (value1 UUID, value2 UUID) AS $$
DECLARE
    playlists_playlist_id UUID;
    platform_playlists_playlist_id UUID;
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM platform_playlists WHERE platform_specific_id = platform_specific_id_input
    ) THEN
        INSERT INTO playlists (id, name, media_url)
        VALUES (uuid_generate_v4(), new_name, new_media_url)
        RETURNING id INTO playlists_playlist_id;

        INSERT INTO platform_playlists (id, platform_id, platform_specific_id, playlist_id)
        VALUES (uuid_generate_v4(), new_platform_id, platform_specific_id_input, playlists_playlist_id)
        RETURNING id INTO platform_playlists_playlist_id;

        RETURN QUERY SELECT playlists_playlist_id, platform_playlists_playlist_id;
    ELSE
        SELECT playlists.id FROM platform_playlists
        JOIN playlists ON platform_playlists.playlist_id = playlists.id
        WHERE platform_specific_id = platform_specific_id_input INTO playlists_playlist_id;

        SELECT platform_playlists.id FROM platform_playlists
        WHERE platform_specific_id = platform_specific_id_input INTO platform_playlists_playlist_id;

        RETURN QUERY SELECT playlists_playlist_id, platform_playlists_playlist_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

DROP FUNCTION IF EXISTS insert_new_song;
CREATE OR REPLACE FUNCTION insert_new_song(new_platform_id VARCHAR(1024), album_id UUID, platform_specific_song_id VARCHAR(1024), new_name VARCHAR(1024), new_media_url VARCHAR(1024), new_preview_url VARCHAR(1024))
RETURNS TABLE (value1 UUID, value2 UUID) AS $$
DECLARE
    songs_song_id UUID;
    platform_songs_song_id UUID;
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM platform_songs WHERE platform_specific_id = platform_specific_song_id
    ) THEN
        INSERT INTO songs (id, name, album_id, media_url, preview_url)
        VALUES (uuid_generate_v4(), new_name, album_id, new_media_url, new_preview_url)
        RETURNING id INTO songs_song_id;

        INSERT INTO platform_songs (id, platform_id, platform_specific_id, song_id)
        VALUES (uuid_generate_v4(), new_platform_id, platform_specific_song_id, songs_song_id)
        RETURNING id INTO platform_songs_song_id;

        RETURN QUERY SELECT songs_song_id, platform_songs_song_id;
    ELSE
        SELECT songs.id FROM platform_songs
        JOIN songs ON platform_songs.song_id = songs.id
        WHERE platform_specific_id = platform_specific_song_id INTO songs_song_id;

        SELECT platform_songs.id FROM platform_songs
        WHERE platform_specific_id = platform_specific_song_id INTO platform_songs_song_id;

        RETURN QUERY SELECT songs_song_id, platform_songs_song_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

INSERT INTO images (id, path) VALUES ('7c5c5c3f-6319-4559-82ba-a52106dac824', 'static/placeholder.png');

INSERT INTO platforms (id, name, icon_id) VALUES ('spotify', 'Spotify', '7c5c5c3f-6319-4559-82ba-a52106dac824');
INSERT INTO platforms (id, name, icon_id) VALUES ('deezer', 'Deezer', '7c5c5c3f-6319-4559-82ba-a52106dac824');

INSERT INTO users VALUES ('6dc10487-60c6-41f8-a2fd-7a450bc3db2a', 'email', 'username', '$argon2id$v=19$m=65536,t=1,p=24$q1OaktL8qTaXZ2M3gi+Z8Q$HYUty9gm/BH1CQc+tQ2+Yc6nUWpsAKXTIxrdRdbcC7A');
INSERT INTO sessions VALUES
('3df14b9a-70c8-4a00-aff1-185d132be749', '6dc10487-60c6-41f8-a2fd-7a450bc3db2a', CURRENT_TIMESTAMP + INTERVAL '1 day');

INSERT INTO connections VALUES ('13ca08cf-3141-40b5-820f-736000432976', '6dc10487-60c6-41f8-a2fd-7a450bc3db2a', '', '', CURRENT_TIMESTAMP);

INSERT INTO users (id, email, username, password_hash) VALUES
  ('8303997f-b12c-4c2b-af9a-7ebe22d5c051', 'user1@example.com', 'user1', 'hash1'),
  ('aa0bfc16-a067-46f5-8821-839a9f01564c', 'user2@example.com', 'user2', 'hash2'),
  ('53336c2a-6985-430e-968d-fae2a921ba9f', 'user3@example.com', 'user3', 'hash3');
