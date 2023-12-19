export default class PostModel{
    id: string = 'id'
    user_id: string = 'user_id'
    caption: string = 'caption' 
    type: string = 'type'
    type_id: string = 'type_id'

    public PostModel(id: string, user_id: string, caption: string, type: string, type_id: string){
        this.id = id
        this.user_id = user_id
        this.caption = caption
        this.type = type
        this.type_id = type_id
    }
}