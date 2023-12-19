package scanner

import (
    "backend/internal/repositories"
    "context"
    "fmt"
    "slices"

    "github.com/google/uuid"
    "github.com/zmb3/spotify"
)

func ScanSpotify(client spotify.Client, connectionId uuid.UUID) {
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
    // fmt.Printf("Fetching Playlists: ")
    // fetchedPlaylists := FetchSpotifyUserPlaylists(&client)
    // missingPlaylists := CheckSpotifyPlaylists(&fetchedPlaylists, libraryId)
    // fmt.Printf("%d playlists\n\rFetching Playlist Tracks: ", len(missingPlaylists))
    // missingPlaylistTracks := FetchSpotifyPlaylistTracks(&client, &missingPlaylists)
    // missingPlaylistTracks = uniqueTracks(missingPlaylistTracks)
    // fmt.Printf("%d tracks\n\rFetching Albums: ", len(missingPlaylistTracks))
    // fetchedAlbums := FetchSpotifyUserAlbums(&client)
    // missingAlbums := CheckSpotifyAlbums(&fetchedAlbums)
    // fmt.Printf("%d albums\n\rFetching Album Tracks: ", len(missingAlbums))
    // missingAlbumTracks := FetchSpotifyAlbumTracks(&client, &missingAlbums)
    // fmt.Printf("%d tracks\n\rFetching Tracks: ", len(missingAlbumTracks))
    // fetchedTracks := FetchSpotifyUserTracks(&client)
    // fetchedTracks = append(fetchedTracks, missingPlaylistTracks...)
    // fetchedTracks = append(fetchedTracks, missingAlbumTracks...)
    // missingTracks := CheckSpotifyTracks(&fetchedTracks)
    // SaveTracks(&missingTracks)
    // fmt.Printf("%d tracks\n\rFetching Artists: ", len(missingTracks))
    // fetchedArtists := FetchSpotifyUserArtists(&client)
    // fmt.Printf("%d artists\n\r", len(fetchedArtists))
    // SaveSpotifyArtists(&fetchedArtists)
    // fmt.Printf("Done")
    // TODO: Fetch Albums
    // TODO: Fetch Artists
    // TODO: Fetch Playlists
    // TODO: Fetch Songs
    // TODO: Fetch Playlist Songs
    // TODO: Add Playlist Songs and Songs Uniquely
    // TODO: Fetch Song Artists and Albums
    // TODO: Add Song Albums and Albums
    // TODO: Add Song Artists and Artists
    // TODO: Fetch Album Songs
    // TODO: Fetch Album Artists
    // TODO: Add Album Songs and Songs Uniquely
    // TODO: Add Combined Artists and Album Artists
    // TODO: Save Artists
    // TODO: Save Albums
    // TODO: Save Songs
    // TODO: Save Playlists
    // TODO: Save Songs Playlist Associations
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

func FetchSpotifyUserPlaylists(client *spotify.Client) []spotify.SimplePlaylist {
    playlistsPage, err := client.CurrentUsersPlaylists()
    if err != nil {
        fmt.Printf("err current playlists")
        return nil
    }
    playlists := playlistsPage.Playlists
    for true {
        err := client.NextPage(playlistsPage)
        if err == spotify.ErrNoMorePages {
            break
        }
        playlists = append(playlists, playlistsPage.Playlists...)
    }
    return playlists
}

func FetchSpotifyUserAlbums(client *spotify.Client) []spotify.FullAlbum {
    savedAlbumsPage, err := client.CurrentUsersAlbums()
    if err != nil {
        fmt.Printf("err current playlists")
        return nil
    }
    var albums []spotify.FullAlbum
    for _, album := range savedAlbumsPage.Albums {
        albums = append(albums, album.FullAlbum)
    }
    for true {
        err := client.NextPage(savedAlbumsPage)
        if err == spotify.ErrNoMorePages {
            break
        }
        for _, album := range savedAlbumsPage.Albums {
            albums = append(albums, album.FullAlbum)
        }
    }
    return albums
}

func FetchSpotifyUserTracks(client *spotify.Client) []spotify.FullTrack {
    savedTracksPage, err := client.CurrentUsersTracks()
    if err != nil {
        fmt.Printf("err current playlists")
        return nil
    }
    var tracks []spotify.FullTrack
    for _, track := range savedTracksPage.Tracks {
        tracks = append(tracks, track.FullTrack)
    }
    for true {
        err := client.NextPage(savedTracksPage)
        if err == spotify.ErrNoMorePages {
            break
        }
        for _, track := range savedTracksPage.Tracks {
            tracks = append(tracks, track.FullTrack)
        }
    }
    return tracks
}

func FetchSpotifyUserArtists(client *spotify.Client) []spotify.FullArtist {
    after := "-1"
    limit := 5
    returned := limit
    var artists []spotify.FullArtist
    for returned == limit {
        list, err := client.CurrentUsersFollowedArtistsOpt(limit, after)
        if err != nil {
            fmt.Printf("%s", err)
            break
        }
        returned = len(list.Artists)
        artists = append(artists, list.Artists...)
        after = list.Artists[len(list.Artists)-1].ID.String()
    }
    return artists
}

func CheckSpotifyPlaylists(playlists *[]spotify.SimplePlaylist, libraryId uuid.UUID) []spotify.SimplePlaylist {
    var ids []string
    for _, playlist := range(*playlists) {
        var id string
        sqlStatement := `
            SELECT insert_new_playlist(CAST ($1 AS UUID), $2, $3);
        `
        err := repositories.Pool.QueryRow(context.Background(),
            sqlStatement,
            libraryId,
            playlist.ID.String(),
            playlist.Name,
        ).Scan(&id)
        if err != nil {
            fmt.Printf("error: %v", err)
        }
        ids = append(ids, id)
    }
    var missing []spotify.SimplePlaylist
    for _, playlist := range(*playlists) {
        if slices.Contains(ids, playlist.ID.String()) {
            missing = append(missing, playlist)
        }
    }
    return missing
}

func FetchSpotifyPlaylistTracks(client *spotify.Client, playlists *[]spotify.SimplePlaylist) []spotify.FullTrack {
    var tracks []spotify.FullTrack
    for i, playlist := range *playlists {
        fmt.Printf("Playlist: %d/%d\n\r", i, len(*playlists))
        playlistTracks, err := client.GetPlaylistTracks(playlist.ID)
        if err != nil {
            fmt.Printf("%v", err)
        }
        for _, track := range playlistTracks.Tracks {
            tracks = append(tracks, track.Track)
        }
        break;
    }
    return tracks
}

func CheckSpotifyAlbums(albums *[]spotify.FullAlbum) []spotify.FullAlbum {
    return *albums
}

func FetchSpotifyAlbumTracks(client *spotify.Client, albums *[]spotify.FullAlbum) []spotify.FullTrack {
    var simpleTracks []spotify.ID
    for _, album := range *albums {
        albumTracks, err := client.GetAlbumTracks(album.ID)
        if err != nil {
            fmt.Printf("%v", err)
        }
        for _, simpleTrack := range *&albumTracks.Tracks {
            simpleTracks = append(simpleTracks, simpleTrack.ID)
        }
    }
    if len(simpleTracks) == 0 {
        return make([]spotify.FullTrack, 0)
    }
    tracks, err := client.GetTracks(simpleTracks...)
    if err != nil {
        fmt.Printf("%v", err)
    }
    tracksSlice := make([]spotify.FullTrack, len(tracks))
    for i, pointer := range tracks {
        tracksSlice[i] = *pointer
    }
    return tracksSlice
}

func CheckSpotifyTracks(tracks *[]spotify.FullTrack) []spotify.FullTrack {
    return *tracks
}

func SaveTracks(tracks *[]spotify.FullTrack) {

}

func SaveSpotifyArtists(artists *[]spotify.FullArtist) {
}
