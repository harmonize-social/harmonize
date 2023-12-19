import type ArtistModel from "./artist";
import type SongModel from "./song";
import type Song from "./song";
export default class AlbumModel{
    id: string = 'id'
    title: string = 'title'
    artists: ArtistModel[] = []
    songs: SongModel[] = []
    url: string = 'url'

    public AlbumModel(id: string, title: string, artists: ArtistModel[], songs: SongModel[], url: string){
        this.id = id
        this.title = title
        this.artists = artists
        this.songs = songs
        this.url = url
    }
}