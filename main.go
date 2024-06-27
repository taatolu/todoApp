package main

import (
    "main/models"
    
    "fmt"
    "log"

    )



func main(){
    
    //log.Println("this is a TEST")

    u := &models.User{}
    u.Name = "test2"
    u.Email = "test2@exsample.com"
    u.Password = "testtest"
    
    u.CreateUser()
    
    
    addUser, err := models.GetUser(2)
    if err !=nil{
        log.Fatalln(err)
    }
    fmt.Println(addUser)
    
    addUser.Name = "yuza"
    addUser.Email = "yuza@example.com"
    addUser.UpdateUser()
    
    updatedUser, err := models.GetUser(2)
    if err !=nil{
        log.Fatalln(err)
    }
    fmt.Println(updatedUser)
    
    _ = models.DeleteUser(2)
    
}