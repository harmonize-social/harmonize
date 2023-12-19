import type song from "./song"

export default interface PlaylistModel{
    id: string
    name: string
    songs: song[]
    url: string
}