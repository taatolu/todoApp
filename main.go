package main

import (
    "main/user"
    _ "main/app"
    "main/app"
    
    "fmt"
    "log"

    )



func main(){
    user1 := user.NewUser("A123","Yusaku")
    fmt.Println(user1)
    
    log.Println("this is a TEST")
    
    u := &app.User{}
    u.Name = "test"
    u.Email = "test@exsample.com"
    u.Password = "testtest"
    fmt.Println(u)
    
    u.CreateUser()
    
}