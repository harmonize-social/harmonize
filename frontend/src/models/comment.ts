import PostModel from "./post"
import type SongModel from "./song"
import type AlbumModel from "./album"
import type PlaylistModel from "./playlist"

export default class CommentModel extends PostModel{
    comment_id: string = 'id'
    user_comment_id: string = 'user_id'
    content_comment: string = 'content'

    constructor(data: {id: string, user_id: string, caption: string, type: string, type_id: string, content: SongModel | AlbumModel | PlaylistModel}, comment_id: string, user_comment_id: string, content_comment: string){
        super(data)
        this.comment_id = comment_id
        this.user_comment_id = user_comment_id
        this.content_comment = content_comment
    }
}