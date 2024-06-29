package models

import(
    "log"
    "time"
    )
    

type Todo struct{
    ID  int
    Content string
    UserID  int
    CreateAt    time.Time
    UpdateAt    time.Time
}


func (u *User) CreateTodo(content string)(err error){
    cmd := `insert into todos(
        content,
        userid,
        create_at,
        update_at) values($1,$2,$3,$4)`
    
    _, err = Db.Exec(cmd,
        content,
        u.ID,
        time.Now(),
        time.Now())
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
        &todo.UserID,
        &todo.CreateAt,
        &todo.UpdateAt)
        
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
            &todo.UserID,
            &todo.CreateAt,
            &todo.UpdateAt)
        if err != nil{
            log.Fatalln(err)
        }
        
        todos = append(todos, todo)
    }
    rows.Close()
    
    return todos, err
}