export default class ArtistModel{
    id: string = 'id'
    name: string = 'name'
    url: string = 'url'

    public ArtistModel(id: string, name: string, url: string){
        this.id = id
        this.name = name
        this.url = url
    }
}