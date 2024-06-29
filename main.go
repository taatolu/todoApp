package main

import (
    "main/models"
    
    "fmt"
    "log"

    )



func main(){

    
    //log.Println("this is a TEST")
    /*
    u := &models.User{}
    u.Name = "testman"
    u.Email = "testman@exsample.com"
    u.Password = "testtest"
    
    u.CreateUser()
    */
    user1, err := models.GetUser(1)
    if err != nil{
        log.Fatalln(err)
    }
    
    user1.CreateTodo("User1のテストTODO2")

        
    todos, err := user1.GetTodos()
    if err != nil {
        log.Fatalln(err)
    }
    for _, v := range todos{
        fmt.Println(v)
    }
}