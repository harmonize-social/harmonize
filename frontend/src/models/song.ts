import type AlbumModel from "./album"
import type ArtistModel from "./artist"

export default class SongModel{
    id: string = 'id'
    album: AlbumModel
    artists: ArtistModel[]
    title: string = 'title'
    mediaUrl: string = 'url'
    previewUrl: string = 'previewUrl'

    constructor(id: string, title: string, artists: ArtistModel[], url: string, album: AlbumModel, previewUrl: string){
        this.id = id
        this.title = title
        this.artists = artists
        this.mediaUrl = url
        this.album = album
        this.previewUrl = previewUrl
    }
}
