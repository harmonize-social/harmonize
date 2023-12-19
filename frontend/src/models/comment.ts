export default class CommentModel{
    id: string = 'id'
    user_id: string = 'user_id'
    post_id: string = 'post_id'
    content: string = 'content'

    public CommentModel(id: string, user_id: string, post_id: string, content: string){
        this.id = id
        this.user_id = user_id
        this.post_id = post_id
        this.content = content
    }
}