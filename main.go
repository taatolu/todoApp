package main

import (
    "main/models"
    "main/utils"
    "main/config"
    "fmt"
    "log"
    )



func main(){
    //設定ファイルの読み込み
    conf, err := config.LoadConfig("product")
    if err != nil{
        log.Fatalln(err)
    }
    
    //LOｇの設定
    utils.LoggingSettings(conf.Logfile)
    
    //ＤＢイニシャライズ
    models.InitDB(conf)
    fmt.Println(models.DB)
    
    
    user1, err := models.GetUser(1)
    if err != nil{
        log.Fatalln(err)
    }
    
    if err = user1.CreateTodo("サンプルだよー"); err!= nil{
        log.Fatalln(err)
    }
    
    /*
    todos, err := user1.GetTodos()
    if err != nil{
        log.Fatalln(err)
    }
    
    for _, v :=range todos{
        fmt.Println(v)
    }
    */
    
    todo1, err := models.GetTodo(1)
    if err != nil{
        log.Fatalln(err)
    }
    fmt.Println(todo1)
    
    err = todo1.UpdateTodo("着手")
    if err != nil{
        log.Fatalln(err)
    }
    /*   
    u := &models.User{}
    u.Name = "testman"
    u.Email = "testman@exsample.com"
    u.Password = "testtest"
    
    //u.CreateUser()
    
    
    //user1.CreateTodo("にこめのさんぷるだよー")
    todo, err := models.GetTodo(2)
    if err != nil{
        log.Fatalln(err)
    }
    fmt.Println(todo)
    
    
    todo1, err := models.GetTodo(1)
    if err != nil{
        log.Fatalln(err)
    }
    
    err = todo1.UpdateTodo("Change TODOOOOO")
    if err != nil{
        log.Fatalln(err)
    }
    
    //user1.CreateTodo("テストTODO1")
    
    
    err = models.DeleteTodos(3)
    if err != nil{
        log.Fatalln(err)
    }
    
        
    todos, err := user1.GetTodos()
    if err != nil {
        log.Fatalln(err)
    }
    for _, v := range todos{
        fmt.Println(v)
    }
    */
}