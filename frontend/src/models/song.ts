import type ArtistModel from "./artist"

export default class SongModel{
    id: string = 'id'
    artists: ArtistModel[] = []
    title: string = 'title'
    mediaUrl: string = 'url'
    previewUrl: string = 'previewUrl'

    constructor(id: string, title: string, url: string, artists: ArtistModel[], previewUrl: string){
        this.id = id
        this.title = title
        this.mediaUrl = url
        this.artists = artists
        this.previewUrl = previewUrl
    }
}


  