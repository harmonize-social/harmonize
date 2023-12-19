import type song from "./song"

export default interface Playlist{
    id: string
    name: string
    songs: song[]
    url: string
}