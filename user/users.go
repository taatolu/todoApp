package user

import (
    //"main/autNum"
    )

type User struct {
    Id  string
    FirstName    string
    LastName    string
}


func NewUser(fname string, lname string) *User{
    user := new(User)
    user.FirstName = fname
    user.LastName = lname
    //user.Id = autNum(fname)
    
    return user
}