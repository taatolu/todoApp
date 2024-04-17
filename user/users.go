package user

type User struct {
    Id  string
    Name    string
}


func NewUser(id string,name string) *User{
    user := new(User)
    user.Id = id
    user.Name = name
    return user
}