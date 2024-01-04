import type AlbumModel from "./album"

export default class SongModel{
    id: string = 'id'
    album: AlbumModel
    title: string = 'title'
    mediaUrl: string = 'url'
    previewUrl: string = 'previewUrl'

    constructor(id: string, title: string, url: string, album: AlbumModel, previewUrl: string){
        this.id = id
        this.title = title
        this.mediaUrl = url
        this.album = album
        this.previewUrl = previewUrl
    }
}
