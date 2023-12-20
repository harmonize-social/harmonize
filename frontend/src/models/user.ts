export default class UserModel{
    id: string = 'id'
    name: string = 'name'
    email: string = 'email'
    password: string = 'password'
    constructor(id: string, name: string, email: string, password: string){
        this.id = id
        this.name = name
        this.email = email
        this.password = password
    }
}