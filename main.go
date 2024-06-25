package main

import (
    "main/user"
    _ "main/app"
    "main/app"
    
    "fmt"
    _"log"

    )



func main(){
    user1 := user.NewUser("A123","Yusaku")
    fmt.Println(user1)
    
    //log.Println("this is a TEST")

    u := &app.User{}
    u.Name = "test2"
    u.Email = "test2@exsample.com"
    u.Password = "testtest"
    
    u.CreateUser()
    
    
    addUser, err := app.GetUser(2)
    if err !=nil{
        log.Fatalln(err)
    }
    fmt.Println(addUser)
    
    addUser.Name = "yuza"
    addUser.Email = "yuza@example.com"
    addUser.UpdateUser()
    
    updatedUser, err := app.GetUser(2)
    if err !=nil{
        log.Fatalln(err)
    }
    fmt.Println(updatedUser)
    
    _ = app.DeleteUser(2)
    
}