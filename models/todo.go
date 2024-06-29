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


func GetTodo(todoid int)(todo Todo, err error){
    todo = Todo{}
    cmd := `select * from todos where id = $1`
    err = Db.QueryRow(cmd,todoid).Scan(
        &todo.ID,
        &todo.Content,
        &todo.UserID)
        
    if err != nil{
        log.Fatalln(err)
    }
    
    return todo, err
}


func (u *User)GetTodos()(todos []Todo, err error){
    cmd := `select * from todos where userid = $1`
    
    rows, err := Db.Query(cmd,u.ID)
    if err != nil {
        log.Fatalln(err)
    }
    
    for rows.Next() {
        var todo Todo
        err = rows.Scan(
            &todo.ID,
            &todo.Content,
            &todo.UserID)
        if err != nil{
            log.Fatalln(err)
        }
        
        todos = append(todos, todo)
    }
    rows.Close()
    
    return todos, err
}