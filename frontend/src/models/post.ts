// import type AlbumModel from './album';
// import type PlaylistModel from './playlist';
// import type SongModel from './song';

export default class PostModel {
    id: string = 'id';
    username: string = 'user_name';
    caption: string = 'caption';
    createdAt: string = 'created_at';
    type: string = 'type';
    content: Object;
    likeCount: number;
    hasLiked: boolean;
    hasSaved: boolean;

    constructor(data: {
        id: string;
        user_name: string;
        createdAt:string;
        caption: string;
        type: string;
        content: Object;
        likes: number;
        liked: boolean;
        saved: boolean;
    }) {
        this.id = data.id;
        this.username = data.user_name;
        this.caption = data.caption;
        this.createdAt = data.createdAt;
        this.type = data.type;
        this.content = data.content;
        this.likeCount = data.likes;
        this.hasLiked = data.liked;
        this.hasSaved = data.saved;
    }
}
