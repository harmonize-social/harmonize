package scanner

import (
    "backend/internal/repositories"
    "context"
    "fmt"

    "github.com/google/uuid"
    "github.com/schollz/progressbar/v3"
    "github.com/zmb3/spotify"
    "go.uber.org/ratelimit"
)

func Spotify(client spotify.Client, connectionId uuid.UUID) {
    fmt.Printf("Create library\n\r")
    sqlStatement := `
    INSERT INTO libraries (platform_id, id, connection_id) VALUES ('spotify', uuid_generate_v4(), $1) RETURNING id;
    `
    var libraryId uuid.UUID
    err := repositories.Pool.QueryRow(context.Background(),
        sqlStatement,
        connectionId).Scan(&libraryId)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    // TODO: Fetch Albums
    standalone_albums, err := FetchUserAlbums(&client)
    if err != nil {
        fmt.Printf("err in fetching albums: %v", err)
        return
    }
    fmt.Println("albums: ", len(standalone_albums))
    // TODO: Fetch Artists
    standalone_artists, err := FetchUserArtists(&client)
    if err != nil {
        fmt.Printf("err in fetching artists: %v", err)
        return
    }
    // TODO: convert artists to simple artists
    standalone_simple_artists := make([]spotify.SimpleArtist, 0)
    for _, artist := range standalone_artists {
        standalone_simple_artists = append(standalone_simple_artists, artist.SimpleArtist)
    }
    // TODO: Fetch Playlists
    fmt.Println("fetching playlists")
    standalone_playlists, err := FetchUserPlaylists(&client)
    if err != nil {
        fmt.Printf("err in fetching playlists: %v", err)
        return
    }
    fmt.Println("fetched playlists: ", len(standalone_playlists))
    // TODO: Fetch Songs
    fmt.Println("fetching tracks")
    standalone_tracks, err := FetchUserTracks(&client)
    if err != nil {
        fmt.Printf("err in fetching tracks: %v", err)
        return
    }
    fmt.Println("fetched tracks: ", len(standalone_tracks))
    // TODO: Fetch Playlist Songs
    fmt.Println("fetching playlist tracks for ", len(standalone_playlists), " playlists")
    playlist_tracks, err := FetchPlaylistTracks(&client, &standalone_playlists)
    if err != nil {
        fmt.Printf("err in fetching playlist tracks: %v", err)
        return
    }
    fmt.Println("fetched playlist tracks: ", len(playlist_tracks))
    // TODO: Add Playlist Songs and Songs
    all_tracks := make([]spotify.FullTrack, 0)
    for _, tracks := range playlist_tracks {
        all_tracks = append(all_tracks, tracks...)
    }
    all_tracks = append(all_tracks, standalone_tracks...)
    // TODO: Make unique
    all_tracks = uniqueTracks(all_tracks)
    // TODO: Add albums
    all_simple_track_albums := make([]spotify.SimpleAlbum, 0)
    for _, track := range all_tracks {
        all_simple_track_albums = append(all_simple_track_albums, track.Album)
    }
    // TODO: Fetch Album Songs
    fmt.Println("fetching: ", len(all_simple_track_albums))
    all_track_albums, err := FetchAlbums(&client, &all_simple_track_albums)
    if err != nil {
        fmt.Printf("err in fetching album tracks: %v", err)
        return
    }
    fmt.Println("fetched Albums: ", len(all_simple_track_albums))
    // TODO: Add Song Albums and Albums
    all_albums := make([]spotify.FullAlbum, 0)
    all_albums = append(all_albums, standalone_albums...)
    all_albums = append(all_albums, all_track_albums...)
    // TODO: Save Albums
    err = SaveAlbums(libraryId, &all_albums)
    if err != nil {
        fmt.Printf("err in saving albums: %v", err)
        return
    }
    // TODO: Save Artists
    err = SaveArtists(libraryId, &standalone_simple_artists)
    if err != nil {
        fmt.Printf("err in saving artists: %v", err)
        return
    }
    // TODO: Save Playlists
    err = SavePlaylists(libraryId, &standalone_playlists)
    if err != nil {
        fmt.Printf("err in saving playlists: %v", err)
    }
    err = SavePlaylistTracks(libraryId, &playlist_tracks)
    if err != nil {
        fmt.Printf("err in saving playlist tracks: %v", err)
    }
}

func uniqueTracks(slice []spotify.FullTrack) []spotify.FullTrack {
    encountered := map[string]bool{}
    unique := []spotify.FullTrack{}

    for _, item := range slice {
        id := item.ID.String()
        if !encountered[id] {
            encountered[id] = true
            unique = append(unique, item)
        }
    }
    return unique
}

func uniqueArtists(slice []spotify.SimpleArtist) []spotify.SimpleArtist {
    encountered := map[string]bool{}
    unique := []spotify.SimpleArtist{}

    for _, item := range slice {
        id := item.ID.String()
        if !encountered[id] {
            encountered[id] = true
            unique = append(unique, item)
        }
    }
    return unique
}

func FetchUserPlaylists(client *spotify.Client) ([]spotify.SimplePlaylist, error) {
    rl := ratelimit.New(2)
    rl.Take()
    playlistsPage, err := client.CurrentUsersPlaylists()
    if err != nil {
        return nil, err
    }
    playlists := playlistsPage.Playlists
    for true {
        rl.Take()
        err := client.NextPage(playlistsPage)
        if err == spotify.ErrNoMorePages {
            break
        }
        playlists = append(playlists, playlistsPage.Playlists...)
    }
    return playlists, nil
}

func FetchUserAlbums(client *spotify.Client) ([]spotify.FullAlbum, error) {
    rl := ratelimit.New(2)
    rl.Take()
    savedAlbumsPage, err := client.CurrentUsersAlbums()
    if err != nil {
        return nil, err
    }
    var albums []spotify.FullAlbum
    for _, album := range savedAlbumsPage.Albums {
        albums = append(albums, album.FullAlbum)
    }
    for true {
        rl.Take()
        err := client.NextPage(savedAlbumsPage)
        if err == spotify.ErrNoMorePages {
            break
        }
        for _, album := range savedAlbumsPage.Albums {
            albums = append(albums, album.FullAlbum)
        }
    }
    return albums, nil
}

func FetchUserTracks(client *spotify.Client) ([]spotify.FullTrack, error) {
    rl := ratelimit.New(2)
    rl.Take()
    savedTracksPage, err := client.CurrentUsersTracks()
    if err != nil {
        return nil, err
    }
    var tracks []spotify.FullTrack
    for _, track := range savedTracksPage.Tracks {
        tracks = append(tracks, track.FullTrack)
    }
    for true {
        rl.Take()
        err := client.NextPage(savedTracksPage)
        if err == spotify.ErrNoMorePages {
            break
        }
        for _, track := range savedTracksPage.Tracks {
            tracks = append(tracks, track.FullTrack)
        }
    }
    return tracks, nil
}

func FetchUserArtists(client *spotify.Client) ([]spotify.FullArtist, error) {
    rl := ratelimit.New(2)
    after := "-1"
    limit := 5
    returned := limit
    var artists []spotify.FullArtist
    for returned == limit {
        rl.Take()
        list, err := client.CurrentUsersFollowedArtistsOpt(limit, after)
        if err != nil {
            return nil, err
        }
        returned = len(list.Artists)
        artists = append(artists, list.Artists...)
        after = list.Artists[len(list.Artists)-1].ID.String()
    }
    return artists, nil
}

func FetchPlaylistTracks(client *spotify.Client, playlists *[]spotify.SimplePlaylist) (map[string][]spotify.FullTrack, error) {
    rl := ratelimit.New(2)
    playlistMap := make(map[string][]spotify.FullTrack)
    bar := progressbar.Default(int64(len(*playlists)))
    for _, playlist := range *playlists {
        bar.Add(1)
        var tracks []spotify.FullTrack
        rl.Take()
        playlistTracks, err := client.GetPlaylistTracks(playlist.ID)
        if err != nil {
            bar.Exit()
            return nil, err
        }
        for _, track := range playlistTracks.Tracks {
            tracks = append(tracks, track.Track)
        }
        playlistMap[playlist.ID.String()] = tracks
    }
    bar.Finish()
    return playlistMap, nil
}

func FetchAlbums(client *spotify.Client, albums *[]spotify.SimpleAlbum) ([]spotify.FullAlbum, error) {
    rl := ratelimit.New(2)
    fullAlbumIds := make([]spotify.ID, len(*albums))
    for _, album := range *albums {
        fullAlbumIds = append(fullAlbumIds, album.ID)
    }
    albumChunks := chunkBy(fullAlbumIds, 20)
    fullAlbums := make([]*spotify.FullAlbum, 0)
    bar := progressbar.Default(int64(len(albumChunks)))
    for _, chunk := range albumChunks {
        rl.Take()
        bums, err := client.GetAlbums(chunk...)
        if err != nil {
            return nil, err
        }
        bar.Add(1)
        fullAlbums = append(fullAlbums, bums...)
    }
    ownedAlbums := make([]spotify.FullAlbum, 0)
    for _, album := range fullAlbums {
        if album == nil {
            continue
        }
        ownedAlbums = append(ownedAlbums, *album)
    }
    bar.Finish()
    return ownedAlbums, nil
}

func chunkBy[T any](items []T, chunkSize int) [][]T {
    var _chunks = make([][]T, 0, (len(items)/chunkSize)+1)
    for chunkSize < len(items) {
        items, _chunks = items[chunkSize:], append(_chunks, items[0:chunkSize:chunkSize])
    }
    return append(_chunks, items)
}

func SaveAlbums(libraryId uuid.UUID, albums *[]spotify.FullAlbum) error {
    artistStatement := `SELECT insert_new_artist($1, $2, $3);`
    albumStatement := `SELECT insert_new_album($1, $2, $3, $4);`
    songStatement := `SELECT insert_new_song($1, $2, $3, $4);`
    for _, album := range *albums {
        var artistId uuid.UUID
        if len(album.Artists) == 0 {
            println("album '", album.Name, album.ID, "' has no artists")
            continue
        }
        err := repositories.Pool.QueryRow(context.Background(), artistStatement, libraryId, album.Artists[0].ID.String(), album.Artists[0].Name).Scan(&artistId)
        if err != nil {
            return err
        }
        var albumId uuid.UUID
        err = repositories.Pool.QueryRow(context.Background(), albumStatement, libraryId, artistId, album.ID.String(), album.Name).Scan(&albumId)
        if err != nil {
            return err
        }
        for _, track := range album.Tracks.Tracks {
            var songId uuid.UUID
            err = repositories.Pool.QueryRow(context.Background(), songStatement, libraryId, albumId, track.ID.String(), track.Name).Scan(&songId)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func SaveArtists(libraryId uuid.UUID, artists *[]spotify.SimpleArtist) error {
    artistStatement := `SELECT insert_new_artist($1, $2, $3);`
    for _, artist := range *artists {
        var artistId uuid.UUID
        err := repositories.Pool.QueryRow(context.Background(), artistStatement, libraryId, artist.ID.String(), artist.Name).Scan(&artistId)
        if err != nil {
            return err
        }
    }
    return nil
}

func SavePlaylists(libraryId uuid.UUID, playlists *[]spotify.SimplePlaylist) error {
    // playlistStatement := `SELECT insert_new_playlist($1, $2, $3, $4);`
    return nil
}

func SavePlaylistTracks(libraryId uuid.UUID, playlists *map[string][]spotify.FullTrack) error {
    return nil
}
