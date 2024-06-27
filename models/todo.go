package models

import(
    "log"
    )
    

type Todo struct{
    ID  int
    Content string
    UserID  int
}


func (u *User) CreateTodo(content string)(err error){
    cmd := `insert into todos(
        content,
        userid) values($1,$2)`
    
    _, err = Db.Exec(cmd,content,u.ID)
    if err != nil{
        log.Fatalln(err)
    }
    
    return err
}