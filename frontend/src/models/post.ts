import type AlbumModel from "./album"
import type PlaylistModel from "./playlist"
import type SongModel from "./song"

export default class PostModel {
    id: string = 'id'
    user_id: string = 'user_id'
    caption: string = 'caption' 
    type: string = 'type'
    type_id: string = 'type_id'
    content: SongModel | AlbumModel | PlaylistModel

    constructor(data: {id: string, user_id: string, caption: string, type: string, type_id: string, content: SongModel | AlbumModel | PlaylistModel}){
        this.id = data.id
        this.user_id = data.user_id
        this.caption = data.caption
        this.type = data.type
        this.type_id = data.type_id
        this.content = data.content
    }
}