package handlers

import (
    "backend/internal/models"
    "backend/internal/repositories"
    "context"
    "net/http"
    "strconv"

    "github.com/google/uuid"
    "github.com/gorilla/mux"
    "github.com/zmb3/spotify/v2"
    spotifyauth "github.com/zmb3/spotify/v2/auth"
    "golang.org/x/oauth2"
)

func SongsHandler(w http.ResponseWriter, r *http.Request) {

    setCommonHeaders(w)
    setAdditionalHeaders(w, "GET")

    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }
    id := uuid.MustParse(r.Header.Get("id"))
    user, err := getUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    client, err := SpotifyClientFromRequest(r, &user.ID)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
    }
    tracks, err := client.CurrentUsersTracks(context.Background(), spotify.Limit(limit), spotify.Offset(offset))

    independentTracks := make([]models.Song, len(tracks.Tracks))
    for i, track := range tracks.Tracks {
        artists := make([]models.Artist, len(track.Artists))
        for j, artist := range track.Artists {
            artists[j] = models.Artist{
                ID:   uuid.New(),
                Name: artist.Name,
            }
        }
        independentTracks[i] = models.Song{
            ID:       uuid.New(),
            Title:    track.Name,
            Artists:  artists,
            MediaURL: track.Album.Images[0].URL,
        }
    }

    models.Result(w, independentTracks)

    /*
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

    */
    // models.Result(w, param)
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
    setCommonHeaders(w)
    setAdditionalHeaders(w, "GET")

    limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
    if err != nil {
        limit = 10
    }
    offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
    if err != nil {
        offset = 0
    }
    id := uuid.MustParse(r.Header.Get("id"))
    user, err := getUserFromSession(id)
    if err != nil {
        models.Error(w, http.StatusUnauthorized, "Malformed session")
        return
    }
    client, err := SpotifyClientFromRequest(r, &user.ID)
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
    }
    artists, err := client.CurrentUsersTopArtists(context.Background(), spotify.Limit(limit), spotify.Offset(offset))
    if err != nil {
        models.Error(w, http.StatusInternalServerError, "Try logging into service again")
    }

    independentArtists := make([]models.Artist, len(artists.Artists))
    for i, artist := range artists.Artists {
        independentArtists[i] = models.Artist{
            ID:   uuid.New(),
            Name: artist.Name,
            MediaURL: artist.Images[0].URL,
        }
    }

    models.Result(w, independentArtists)
}

func SpotifyClientFromRequest(r *http.Request, userId *uuid.UUID) (*spotify.Client, error) {
    params := mux.Vars(r)
    param := params["service"]
    var token oauth2.Token
    err := repositories.Pool.QueryRow(r.Context(), "SELECT access_token, refresh_token, expiry FROM connections JOIN libraries ON connections.id = libraries.connection_id WHERE user_id = $1 AND platform_id = $2", userId, param).Scan(&token.AccessToken, &token.RefreshToken, &token.Expiry)
    if err != nil {
        return nil, err
    }
    auth := spotifyauth.New(
        spotifyauth.WithClientID("8c3d77ea95764c898afa8ed598c1db01"),
        spotifyauth.WithClientSecret("3773a05af958442cb77482f1e4601299"),
    )
    newToken, err := auth.RefreshToken(context.Background(), &token)
    if err != nil {
        return nil, err
    }
    newAuth := spotifyauth.New()
    httpClient := newAuth.Client(context.Background(), newToken)
    client := spotify.New(httpClient,
        spotify.WithRetry(true),
    )
    return client, nil
}
