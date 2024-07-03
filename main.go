package main

import (
    "main/models"
    
    "fmt"
    "log"

    )



func main(){


    u := &models.User{}
    u.Name = "testman"
    u.Email = "testman@exsample.com"
    u.Password = "testtest"
    
    //u.CreateUser()
    
    user1, err := models.GetUser(1)
    if err != nil{
        log.Fatalln(err)
    }
    
    
    todo1, err := models.GetTodo(1)
    if err != nil{
        log.Fatalln(err)
    }
    
    err = todo1.UpdateTodo("Change TODOOOOO")
    if err != nil{
        log.Fatalln(err)
    }
    
    //user1.CreateTodo("テストTODO1")
    
    /*
    err = models.DeleteTodos(3)
    if err != nil{
        log.Fatalln(err)
    }
    */
        
    todos, err := user1.GetTodos()
    if err != nil {
        log.Fatalln(err)
    }
    for _, v := range todos{
        fmt.Println(v)
    }
}