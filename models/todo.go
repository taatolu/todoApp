package models

import(
    "fmt"
    "time"
    )

type Todo struct{
    ID  int
    Content     string
    UserID      int
    Create_at   time.Time
    Update_at   time.Time
}

func (u *User)CreateTodo(content string)error{
    cmd:= `insert into todos (
        content,
        userid,
        create_at,
        update_at) values ($1,$2,$3,$4)`
    if _, err := DB.Exec(cmd, content, u.ID, time.Now(), time.Now()); err!= nil{
        return fmt.Errorf("CreteTodoError %w", err)
    }
    
    return nil
}

func GetTodo(todoNum int)(todo *Todo, err error){
    todo = &Todo{}
    cmd := `select * from todos where id = $1`
    err = DB.QueryRow(cmd,todoNum).Scan(
        &todo.ID,
        &todo.Content,
        &todo.UserID,
        &todo.Create_at,
        &todo.Update_at)
    if err != nil{
        return nil, fmt.Errorf("GetTodoError %w", err)
    }
    
    return todo, nil
}


func (u *User)GetTodos()(todos []*Todo, err error){
    
    cmd := `select * from todos where userid = $1`
    
    rows, err := DB.Query(cmd,u.ID)
    if err != nil{
        return nil, fmt.Errorf("GetTodosError %w", err)
    }
    defer rows.Close()
    
    for rows.Next(){
        todo := &Todo{}
        err = rows.Scan(
            &todo.ID,
            &todo.Content,
            &todo.UserID,
            &todo.Create_at,
            &todo.Update_at)
        if err != nil{
            return nil, fmt.Errorf("GetTodosError %w", err)
        }
        todos = append(todos,todo)
    }
    return todos, nil
}

func DeleteTodo(todoid int)error{
    cmd:=`delete from todos where id=$1`
    if _, err = DB.Exec(cmd, todoid); err != nil{
        return fmt.Errorf("DeleteTodosError %w", err)
    }
    return nil
}

