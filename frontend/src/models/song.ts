export default class SongModel {
    id: string = 'id'
    title: string = 'title'
    url: string = 'url'

    public SongModel(id: string, title: string, url: string){
        this.id = id
        this.title = title
        this.url = url
    }
}