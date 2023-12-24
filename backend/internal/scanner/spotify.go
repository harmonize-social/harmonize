package scanner

import ()

// func Spotify(client spotify.Client, connectionId uuid.UUID) {
//     fmt.Printf("Create library\n\r")
//
//     // TODO: Fetch Albums
//     standalone_albums, err := FetchUserAlbums(&client)
//     if err != nil {
//         fmt.Printf("err in fetching albums: %v", err)
//         return
//     }
//     fmt.Println("albums: ", len(standalone_albums))
//     // TODO: Fetch Artists
//     standalone_artists, err := FetchUserArtists(&client)
//     if err != nil {
//         fmt.Printf("err in fetching artists: %v", err)
//         return
//     }
//     // TODO: convert artists to simple artists
//     standalone_simple_artists := make([]spotify.SimpleArtist, 0)
//     for _, artist := range standalone_artists {
//         standalone_simple_artists = append(standalone_simple_artists, artist.SimpleArtist)
//     }
//     // TODO: Fetch Playlists
//     fmt.Println("fetching playlists")
//     standalone_playlists, err := FetchUserPlaylists(&client)
//     if err != nil {
//         fmt.Printf("err in fetching playlists: %v", err)
//         return
//     }
//     fmt.Println("fetched playlists: ", len(standalone_playlists))
//     // TODO: Fetch Songs
//     fmt.Println("fetching tracks")
//     standalone_tracks, err := FetchUserTracks(&client)
//     if err != nil {
//         fmt.Printf("err in fetching tracks: %v", err)
//         return
//     }
//     fmt.Println("fetched tracks: ", len(standalone_tracks))
//     // TODO: Fetch Playlist Songs
//     fmt.Println("fetching playlist tracks for ", len(standalone_playlists), " playlists")
//     playlist_tracks, err := FetchPlaylistTracks(&client, &standalone_playlists)
//     if err != nil {
//         fmt.Printf("err in fetching playlist tracks: %v", err)
//         return
//     }
//     fmt.Println("fetched playlist tracks: ", len(playlist_tracks))
//     // TODO: Add Playlist Songs and Songs
//     all_tracks := make([]spotify.FullTrack, 0)
//     for _, tracks := range playlist_tracks {
//         all_tracks = append(all_tracks, tracks...)
//     }
//     all_tracks = append(all_tracks, standalone_tracks...)
//     // TODO: Make unique
//     all_tracks = uniqueTracks(all_tracks)
//     // TODO: Add albums
//     all_simple_track_albums := make([]spotify.SimpleAlbum, 0)
//     for _, track := range all_tracks {
//         all_simple_track_albums = append(all_simple_track_albums, track.Album)
//     }
//     all_simple_track_albums = uniqueAlbums(all_simple_track_albums)
//     // TODO: Fetch Album Songs
//     fmt.Println("fetching albums: ", len(all_simple_track_albums))
//     all_track_albums, err := FetchAlbums(&client, &all_simple_track_albums)
//     if err != nil {
//         fmt.Printf("err in fetching album tracks: %v", err)
//         return
//     }
//     fmt.Println("fetched albums: ", len(all_simple_track_albums))
//     // TODO: Add Song Albums and Albums
//     all_albums := make([]spotify.FullAlbum, 0)
//     all_albums = append(all_albums, standalone_albums...)
//     all_albums = append(all_albums, all_track_albums...)
//     // TODO: Save Albums
//     err = SaveAlbums(libraryId, &all_albums)
//     if err != nil {
//         fmt.Printf("err in saving albums: %v", err)
//         return
//     }
//     // TODO: Save Artists
//     err = SaveArtists(libraryId, &standalone_simple_artists)
//     if err != nil {
//         fmt.Printf("err in saving artists: %v", err)
//         return
//     }
//     // TODO: Save Playlists
//     err = SavePlaylists(libraryId, &standalone_playlists, &playlist_tracks)
//     if err != nil {
//         fmt.Printf("err in saving playlists: %v", err)
//         return
//     }
//     err = SaveTracks(libraryId, &standalone_tracks)
//     if err != nil {
//         fmt.Printf("err in saving tracks: %v", err)
//         return
//     }
// }
//
// func uniqueAlbums(slice []spotify.SimpleAlbum) []spotify.SimpleAlbum {
//     encountered := map[string]bool{}
//     unique := []spotify.SimpleAlbum{}
//
//     for _, item := range slice {
//         id := item.ID.String()
//         if !encountered[id] {
//             encountered[id] = true
//             unique = append(unique, item)
//         }
//     }
//     return unique
// }
//
// func uniqueTracks(slice []spotify.FullTrack) []spotify.FullTrack {
//     encountered := map[string]bool{}
//     unique := []spotify.FullTrack{}
//
//     for _, item := range slice {
//         id := item.ID.String()
//         if !encountered[id] {
//             encountered[id] = true
//             unique = append(unique, item)
//         }
//     }
//     return unique
// }
//
// func uniqueArtists(slice []spotify.SimpleArtist) []spotify.SimpleArtist {
//     encountered := map[string]bool{}
//     unique := []spotify.SimpleArtist{}
//
//     for _, item := range slice {
//         id := item.ID.String()
//         if !encountered[id] {
//             encountered[id] = true
//             unique = append(unique, item)
//         }
//     }
//     return unique
// }
//
// func FetchUserPlaylists(client *spotify.Client) ([]spotify.SimplePlaylist, error) {
//     rl := ratelimit.New(3)
//     rl.Take()
//     playlistsPage, err := client.CurrentUsersPlaylists(context.Background())
//     if err != nil {
//         return nil, err
//     }
//     playlists := playlistsPage.Playlists
//     for true {
//         rl.Take()
//         err := client.NextPage(context.Background(), playlistsPage)
//         if err == spotify.ErrNoMorePages {
//             break
//         }
//         playlists = append(playlists, playlistsPage.Playlists...)
//     }
//     return playlists, nil
// }
//
// func FetchUserAlbums(client *spotify.Client) ([]spotify.FullAlbum, error) {
//     rl := ratelimit.New(3)
//     rl.Take()
//     savedAlbumsPage, err := client.CurrentUsersAlbums(context.Background())
//     if err != nil {
//         return nil, err
//     }
//     var albums []spotify.FullAlbum
//     for _, album := range savedAlbumsPage.Albums {
//         albums = append(albums, album.FullAlbum)
//     }
//     for true {
//         rl.Take()
//         err := client.NextPage(context.Background(), savedAlbumsPage)
//         if err == spotify.ErrNoMorePages {
//             break
//         }
//         for _, album := range savedAlbumsPage.Albums {
//             albums = append(albums, album.FullAlbum)
//         }
//     }
//     return albums, nil
// }
//
// func FetchUserTracks(client *spotify.Client) ([]spotify.FullTrack, error) {
//     rl := ratelimit.New(3)
//     rl.Take()
//     savedTracksPage, err := client.CurrentUsersTracks(context.Background())
//     if err != nil {
//         return nil, err
//     }
//     var tracks []spotify.FullTrack
//     for _, track := range savedTracksPage.Tracks {
//         tracks = append(tracks, track.FullTrack)
//     }
//     for true {
//         rl.Take()
//         err := client.NextPage(context.Background(), savedTracksPage)
//         if err == spotify.ErrNoMorePages {
//             break
//         }
//         for _, track := range savedTracksPage.Tracks {
//             tracks = append(tracks, track.FullTrack)
//         }
//     }
//     return tracks, nil
// }
//
// func FetchUserArtists(client *spotify.Client) ([]spotify.FullArtist, error) {
//     rl := ratelimit.New(3)
//     after := "-1"
//     limit := 5
//     returned := limit
//     var artists []spotify.FullArtist
//     for returned == limit {
//         rl.Take()
//         list, err := client.CurrentUsersFollowedArtists(
//             context.Background(),
//             spotify.Limit(limit),
//             spotify.After(after),
//         )
//         if err != nil {
//             return nil, err
//         }
//         returned = len(list.Artists)
//         artists = append(artists, list.Artists...)
//         after = list.Artists[len(list.Artists)-1].ID.String()
//     }
//     return artists, nil
// }
//
// func FetchPlaylistTracks(client *spotify.Client, playlists *[]spotify.SimplePlaylist) (map[string][]spotify.FullTrack, error) {
//     rl := ratelimit.New(2)
//     playlistMap := make(map[string][]spotify.FullTrack)
//     bar := progressbar.Default(int64(len(*playlists)))
//     for _, playlist := range *playlists {
//         bar.Add(1)
//         rl.Take()
//         playlistTracks, err := client.GetPlaylistItems(context.Background(), playlist.ID)
//         if err != nil {
//             bar.Exit()
//             return nil, err
//         }
//         var tracks []spotify.FullTrack
//         for _, track := range playlistTracks.Items {
//             if track.Track.Track == nil {
//                 continue
//             }
//             tracks = append(tracks, *track.Track.Track)
//         }
//         playlistMap[playlist.ID.String()] = tracks
//     }
//     bar.Finish()
//     return playlistMap, nil
// }
//
// func FetchAlbums(client *spotify.Client, albums *[]spotify.SimpleAlbum) ([]spotify.FullAlbum, error) {
//     rl := ratelimit.New(2)
//     fullAlbumIds := make([]spotify.ID, len(*albums))
//     for _, album := range *albums {
//         fullAlbumIds = append(fullAlbumIds, album.ID)
//     }
//     albumChunks := chunkBy(fullAlbumIds, 20)
//     fullAlbums := make([]*spotify.FullAlbum, 0)
//     bar := progressbar.Default(int64(len(albumChunks)))
//     for _, chunk := range albumChunks {
//         rl.Take()
//         bums, err := client.GetAlbums(context.Background(), chunk, spotify.Limit(10))
//         if err != nil {
//             return nil, err
//         }
//         bar.Add(1)
//         fullAlbums = append(fullAlbums, bums...)
//     }
//     ownedAlbums := make([]spotify.FullAlbum, 0)
//     for _, album := range fullAlbums {
//         if album == nil {
//             continue
//         }
//         ownedAlbums = append(ownedAlbums, *album)
//     }
//     bar.Finish()
//     return ownedAlbums, nil
// }
//
// func chunkBy[T any](items []T, chunkSize int) [][]T {
//     var _chunks = make([][]T, 0, (len(items)/chunkSize)+1)
//     for chunkSize < len(items) {
//         items, _chunks = items[chunkSize:], append(_chunks, items[0:chunkSize:chunkSize])
//     }
//     return append(_chunks, items)
// }
//
// func SaveAlbums(libraryId uuid.UUID, albums *[]spotify.FullAlbum) error {
//     artistStatement := `SELECT * FROM insert_new_artist($1, $2, $3);`
//     albumStatement := `SELECT * FROM insert_new_album($1, $2, $3);`
//     songStatement := `SELECT * FROM insert_new_song($1, $2, $3, $4);`
//     bar := progressbar.Default(int64(len(*albums)))
//     for _, album := range *albums {
//         var albumId uuid.UUID
//         var platformAlbumId uuid.UUID
//         err := repositories.Pool.QueryRow(context.Background(), albumStatement, libraryId, album.ID.String(), album.Name).Scan(&albumId, &platformAlbumId)
//         if err != nil {
//             fmt.Printf("here", err)
//             return err
//         }
//         for _, artist := range album.Artists {
//             var artistId uuid.UUID
//             var platformArtistId uuid.UUID
//             err := repositories.Pool.QueryRow(context.Background(), artistStatement, libraryId, artist.ID.String(), artist.Name).Scan(&artistId, &platformArtistId)
//             if err != nil {
//                 fmt.Println("print 1")
//                 return err
//             }
//             var artistAlbumId uuid.UUID
//             err = repositories.Pool.QueryRow(context.Background(), `INSERT INTO artists_album (id, artist_id, album_id) VALUES (uuid_generate_v4(), $1, $2) RETURNING id;`, artistId, albumId).Scan(&artistAlbumId)
//             if err != nil {
//                 fmt.Println("print 2")
//                 return err
//             }
//         }
//         for _, track := range album.Tracks.Tracks {
//             var songId uuid.UUID
//             var platformSongId uuid.UUID
//             err = repositories.Pool.QueryRow(context.Background(), songStatement, libraryId, albumId, track.ID.String(), track.Name).Scan(&songId, &platformSongId)
//             if err != nil {
//                 fmt.Println("print 3")
//                 return err
//             }
//         }
//         var userLikedAlbumId uuid.UUID
//         err = repositories.Pool.QueryRow(context.Background(), `INSERT INTO user_liked_albums (id, library_id, album_id) VALUES (uuid_generate_v4(), $1, $2) RETURNING id;`, libraryId, platformAlbumId).Scan(&userLikedAlbumId)
//         if err != nil {
//             fmt.Println("print 4", platformAlbumId)
//             return err
//         }
//         bar.Add(1)
//     }
//     bar.Finish()
//     return nil
// }
//
// func SaveArtists(libraryId uuid.UUID, artists *[]spotify.SimpleArtist) error {
//     artistStatement := `SELECT * FROM insert_new_artist($1, $2, $3);`
//     for _, artist := range *artists {
//         var artistId uuid.UUID
//         var platformArtistId uuid.UUID
//         err := repositories.Pool.QueryRow(context.Background(), artistStatement, libraryId, artist.ID.String(), artist.Name).Scan(&artistId, &platformArtistId)
//         if err != nil {
//             return err
//         }
//         err = repositories.Pool.QueryRow(context.Background(), `INSERT INTO user_followed_artists (id, library_id, artist_id) VALUES (uuid_generate_v4(), $1, $2) RETURNING id;`, libraryId, platformArtistId).Scan(&artistId)
//         if err != nil {
//             return err
//         }
//     }
//     return nil
// }
//
// func SavePlaylists(libraryId uuid.UUID, playlists *[]spotify.SimplePlaylist, playlistMap *map[string][]spotify.FullTrack) error {
//     bar := progressbar.Default(int64(len(*playlists)))
//     for _, playlist := range *playlists {
//         var playlistId uuid.UUID
//         var platformPlaylistId uuid.UUID
//         err := repositories.Pool.QueryRow(context.Background(), `SELECT * FROM insert_new_playlist($1, $2, $3);`, libraryId, playlist.ID.String(), playlist.Name).Scan(&playlistId, &platformPlaylistId)
//         if err != nil {
//             fmt.Println("print 1")
//             return err
//         }
//         err = repositories.Pool.QueryRow(context.Background(), `INSERT INTO user_followed_playlists (id, library_id, playlist_id) VALUES (uuid_generate_v4(), $1, $2) RETURNING id;`, libraryId, platformPlaylistId).Scan(&playlistId)
//         if err != nil {
//             fmt.Println("print 2")
//             return err
//         }
//         songs := (*playlistMap)[playlist.ID.String()]
//         for _, song := range songs {
//             var songId uuid.UUID
//             err = repositories.Pool.QueryRow(context.Background(), `SELECT songs.id FROM platform_songs JOIN songs ON platform_songs.song_id = songs.id WHERE platform_specific_id = $1;`, song.ID.String()).Scan(&songId)
//             if err != nil {
//                 fmt.Println("print 3")
//                 return err
//             }
//             err = repositories.Pool.QueryRow(context.Background(), `INSERT INTO playlist_songs (id, playlist_id, song_id) VALUES (uuid_generate_v4(), $1, $2) RETURNING id;`, playlistId, songId).Scan(&songId)
//             if err != nil {
//                 fmt.Println("print 4", songId, playlist.ID, playlistId)
//                 continue
//             }
//             err = nil
//         }
//         bar.Add(1)
//     }
//     bar.Finish()
//     return nil
// }
//
// func SaveTracks(libraryId uuid.UUID, tracks *[]spotify.FullTrack) error {
//     for _, track := range *tracks {
//         var songId uuid.UUID
//         var platformSongId uuid.UUID
//         err := repositories.Pool.QueryRow(context.Background(), `SELECT * FROM insert_new_song($1, $2, $3, $4);`, libraryId, track.Album.ID.String(), track.ID.String(), track.Name).Scan(&songId, &platformSongId)
//         if err != nil {
//             return err
//         }
//         err = repositories.Pool.QueryRow(context.Background(), `INSERT INTO user_liked_songs (id, library_id, song_id) VALUES (uuid_generate_v4(), $1, $2) RETURNING id;`, libraryId, platformSongId).Scan(&songId)
//         if err != nil {
//             return err
//         }
//     }
//     return nil
// }

/*
DROP FUNCTION IF EXISTS insert_new_artist;
CREATE OR REPLACE FUNCTION insert_new_artist(new_library_id UUID, platform_specific_id_input VARCHAR(1024), new_name VARCHAR(1024))
RETURNS UUID AS $$
DECLARE
    artists_artist_id UUID;
    platform_artists_artist_id UUID;
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM platform_artists WHERE platform_specific_id = platform_specific_id_input
    ) THEN
        INSERT INTO artists (id, name)
        VALUES (uuid_generate_v4(), new_name)
        RETURNING id INTO artists_artist_id;

        INSERT INTO platform_artists (id, platform_specific_id, artist_id)
        VALUES (uuid_generate_v4(), platform_specific_id_input, artists_artist_id)
        RETURNING id INTO platform_artists_artist_id;

        RETURN artists_artist_id;
    ELSE
        SELECT artists.id FROM platform_artists
        JOIN artists ON platform_artists.artist_id = artists.id
        WHERE platform_specific_id = platform_specific_id_input INTO artists_artist_id;
        RETURN artists_artist_id;
    END IF;
END;
$$ LANGUAGE plpgsql;


DROP FUNCTION IF EXISTS insert_new_album;
CREATE OR REPLACE FUNCTION insert_new_album(new_library_id UUID, platform_specific_album_id VARCHAR(1024), new_name VARCHAR(1024))
RETURNS UUID AS $$
DECLARE
    albums_album_id UUID;
    platform_albums_album_id UUID;
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM platform_albums WHERE platform_specific_id = platform_specific_album_id
    ) THEN
        INSERT INTO albums (id, name, author_id)
        VALUES (uuid_generate_v4(), new_name, artist_id)
        RETURNING id INTO albums_album_id;

        INSERT INTO platform_albums (id, platform_specific_id, album_id)
        VALUES (uuid_generate_v4(), platform_specific_album_id, albums_album_id)
        RETURNING id INTO platform_albums_album_id;

        RETURN albums_album_id;
    ELSE
        SELECT albums.id FROM platform_albums
        JOIN albums ON platform_albums.album_id = albums.id
        WHERE platform_specific_id = platform_specific_album_id INTO albums_album_id;
        RETURN albums_album_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

DROP FUNCTION IF EXISTS insert_new_playlist;
CREATE OR REPLACE FUNCTION insert_new_playlist(new_library_id UUID, platform_specific_id_input VARCHAR(1024), new_name VARCHAR(1024))
RETURNS UUID AS $$
DECLARE
    playlists_playlist_id UUID;
    platform_playlists_playlist_id UUID;
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM platform_playlists WHERE platform_specific_id = platform_specific_id_input
    ) THEN
        INSERT INTO playlists (id, name)
        VALUES (uuid_generate_v4(), new_name)
        RETURNING id INTO playlists_playlist_id;

        INSERT INTO platform_playlists (id, platform_specific_id, playlist_id)
        VALUES (uuid_generate_v4(), platform_specific_id_input, playlists_playlist_id)
        RETURNING id INTO platform_playlists_playlist_id;

        RETURN playlists_playlist_id;
    ELSE
        SELECT playlists.id FROM platform_playlists
        JOIN playlists ON platform_playlists.playlist_id = playlists.id
        WHERE platform_specific_id = platform_specific_id_input INTO playlists_playlist_id;
        RETURN playlists_playlist_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

DROP FUNCTION IF EXISTS insert_new_song;
CREATE OR REPLACE FUNCTION insert_new_song(new_library_id UUID, album_id UUID, platform_specific_song_id VARCHAR(1024), new_name VARCHAR(1024))
RETURNS UUID AS $$
DECLARE
    songs_song_id UUID;
    platform_songs_song_id UUID;
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM platform_songs WHERE platform_specific_id = platform_specific_song_id
    ) THEN
        INSERT INTO songs (id, name, album_id)
        VALUES (uuid_generate_v4(), new_name, album_id)
        RETURNING id INTO songs_song_id;

        INSERT INTO platform_songs (id, platform_specific_id, song_id)
        VALUES (uuid_generate_v4(), platform_specific_song_id, songs_song_id)
        RETURNING id INTO platform_songs_song_id;

        RETURN songs_song_id;
    ELSE
        SELECT songs.id FROM platform_songs
        JOIN songs ON platform_songs.song_id = songs.id
        WHERE platform_specific_id = platform_specific_song_id INTO songs_song_id;
        RETURN songs_song_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

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

CREATE TABLE IF NOT EXISTS artists(
    id UUID PRIMARY KEY,
    name VARCHAR(1024) NOT NULL
);

CREATE TABLE IF NOT EXISTS albums(
    id UUID PRIMARY KEY,
    name VARCHAR(1024) NOT NULL
);

CREATE TABLE IF NOT EXISTS artists_album(
    id UUID PRIMARY KEY,
    artist_id UUID REFERENCES artists (id) NOT NULL,
    album_id UUID REFERENCES albums (id) NOT NULL
);


CREATE TABLE IF NOT EXISTS songs(
    id UUID PRIMARY KEY,
    name VARCHAR(1024) NOT NULL,
    album_id UUID REFERENCES albums (id) NOT NULL
);

CREATE TABLE IF NOT EXISTS playlists(
    id UUID PRIMARY KEY,
    name VARCHAR(1024) NOT NULL
);


CREATE TABLE IF NOT EXISTS playlist_songs(
    id UUID PRIMARY KEY,
    playlist_id UUID REFERENCES playlists (id) NOT NULL,
    song_id UUID REFERENCES songs (id) NOT NULL
);
*/
