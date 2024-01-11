# Harmonize

## Running

Easiest way is to `docker-compose up`, after you have configured your environment variables.
If using docker, put the correct environment variables in the correct sections of the `docker-compose.yml`

## Database

PostgresSQL

## Backend

Go + Gorilla/mux

### Environment Variables

```.env
POSTGRES_URL="postgres://user:postgres@172.20.0.2:5432/harmonize"
BACKEND_SECRET="94d9f2dee4bb80fea10399a0e10a9a3df874d86b9454da2d6f5a4bb57314329d" # Optional, this is an example
SPOTIFY_SECRET="Your spotify app secret here"
SPOTIFY_CLIENT_ID="Your spotify app client id here"
SPOTIFY_REDIRECT="http://172.20.0.4:5173/redirect/spotify" # FRONTEND_URL + /redirect/spotify
DEEZER_SECRET="Your deezer app secret here"
DEEZER_CLIENT_ID="Your deezer app client id here"
DEEZER_REDIRECT="http://172.20.0.4:5173/redirect/deezer" # FRONTEND_URL + /redirect/deezer
FRONTEND_URL="http://172.20.0.4:5173" #  Where the frontend is hosted
```

## Frontend

SvelteKit + Svelte

### Environment Variables

```.env
API_URL: "http://172.20.0.3:8080" # Where the frontend can find the backend
```
