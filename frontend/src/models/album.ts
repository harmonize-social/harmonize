import type Artist from "./artist";
import type Song from "./song";
export default interface AlbumModel{
    id: string
    title: string
    artists: Artist[]
    songs: Song[]
    url: string
}