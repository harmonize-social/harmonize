package scanner

import (
    "fmt"

    "github.com/zmb3/spotify"
)

func ScanSpotify(client spotify.Client) {
    fmt.Printf("Fetching Playlists: ")
    fetchedPlaylists := FetchSpotifyUserPlaylists(&client)
    missingPlaylists := CheckSpotifyPlaylists(&fetchedPlaylists)
    fmt.Printf("%d playlists\n\rFetching Playlist Tracks: " , len(missingPlaylists))
    missingPlaylistIds := uniquePlaylists(missingPlaylists)
    missingPlaylistTracks := FetchSpotifyPlaylistTracks(&client, &missingPlaylistIds)
    missingPlaylistTracks = uniqueTracks(missingPlaylistTracks)
    fmt.Printf("%d tracks\n\rFetching Albums: " , len(missingPlaylistTracks))
    fetchedAlbums := FetchSpotifyUserAlbums(&client)
    missingAlbums := CheckSpotifyAlbums(&fetchedAlbums)
    fmt.Printf("%d albums\n\rFetching Album Tracks: " , len(missingAlbums))
    missingAlbumTracks := FetchSpotifyAlbumTracks(&client, &missingAlbums)
    fmt.Printf("%d tracks\n\rFetching Tracks: " , len(missingAlbumTracks))
    fetchedTracks := FetchSpotifyUserTracks(&client)
    fetchedTracks = append(fetchedTracks, missingPlaylistTracks...)
    fetchedTracks = append(fetchedTracks, missingAlbumTracks...)
    missingTracks := CheckSpotifyTracks(&fetchedTracks)
    SaveTracks(&missingTracks)
    fmt.Printf("%d tracks\n\rFetching Artists: " , len(missingTracks))
    fetchedArtists := FetchSpotifyUserArtists(&client)
    fmt.Printf("%d artists\n\r", len(fetchedArtists))
    SaveSpotifyArtists(&fetchedArtists)
    fmt.Printf("Done")
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

func uniquePlaylists(slice []spotify.SimplePlaylist) []string {
    encountered := map[string]bool{}
    unique := []string{}

    for _, item := range slice {
        id := item.ID.String()
        if !encountered[id] {
            encountered[id] = true
            unique = append(unique, id)
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

func CheckSpotifyPlaylists(playlists *[]spotify.SimplePlaylist) []spotify.SimplePlaylist {
    return *playlists
}

func FetchSpotifyPlaylistTracks(client *spotify.Client, playlists *[]string) []spotify.FullTrack {
    var tracks []spotify.FullTrack
    for i, id := range *playlists {
        fmt.Printf("Playlist: %d/%d\n\r", i, len(*playlists))
        playlistTracks, err := client.GetPlaylistTracks(spotify.ID(id))
        if err != nil {
            fmt.Printf("%v", err)
        }
        for _, track := range playlistTracks.Tracks {
            tracks = append(tracks, track.Track)
        }
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
