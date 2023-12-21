import type AlbumModel from './album';
import type PlaylistModel from './playlist';
import type SongModel from './song';

export default class PostModel {
	id: string = 'id';
	user_id: string = 'user_id';
	user_name: string = 'user_name';
	caption: string = 'caption';
	type: string = 'type';
	content: SongModel | AlbumModel | PlaylistModel;
	likes: number;
	liked: boolean = false;
	saved: boolean = false;

	constructor(data: {
		id: string;
		user_id: string;
		user_name: string;
		caption: string;
		type: string;
		content: SongModel | AlbumModel | PlaylistModel;
		likes: number;
		liked: boolean;
		saved: boolean;
	}) {
		this.id = data.id;
		this.user_id = data.user_id;
		this.user_name = data.user_name;
		this.caption = data.caption;
		this.type = data.type;
		this.content = data.content;
		this.likes = data.likes;
		this.liked = data.liked;
		this.saved = data.saved;
	}
}
