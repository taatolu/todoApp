package models

import(
    "testing"
    _ "github.com/stretchr/testify/assert"
    "fmt"
    )


func TestCreateTodo(t *testing.T){
    //まずはTodoを作成するためのUser作成
    sampleUser := &User{
        Name:   "sample",
        Email:  "sample@exam.com",
        Password:   "testtest",
    }
    if err :=  sampleUser.CreateUser(); err !=nil{
        t.Errorf("Todoを作成するためのUser作成で失敗 %w", err)
    }
    //エラーが無かったら、CreateUserメソッドからreturningでIDがとれる
    fmt.Println(sampleUser.ID)
}