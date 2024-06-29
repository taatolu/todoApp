package main

import (
    "main/models"
    
    "fmt"
    "log"

    )



func main(){
    
    u, err:= models.GetUser(2)
    if err != nil{
        log.Fatalln(err)
    }
    
    todos, err := u.GetTodos()
    if err != nil {
        log.Fatalln(err)
    }
    for _, v := range todos{
        fmt.Println(v)
    }
    
    //log.Println("this is a TEST")
    /*
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
    
    updatedUser.CreateTodo("テストTODO")
    fmt.Println(updatedUser)
    
    _ = models.DeleteUser(2)
    */
}