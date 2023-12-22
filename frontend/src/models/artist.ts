export default class ArtistModel{
    id: string = 'id'
    name: string = 'name'
    url: string = 'url'
	isFollowing: any;
	artistAlbums: any;
	popularSongs: any;
	alt: any;
	image: any;

   constructor(id: string, name: string, url: string){
        this.id = id
        this.name = name
        this.url = url
    }
}