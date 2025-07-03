package models

import(
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
    )


func TestCreateTodo(t *testing.T){
    //Todo作成のためのUser作成
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
    //Todo作成のためのUser（上で作成したUser）を取得
    user, err := GetUser(sampleUser.ID)
    assert.NoError(t, err, "Todo作成のためのUser取得失敗")
    
    
    //tableテストのための構造体作成
    tests := []struct{
        testname    string
        content     string
        wantError   bool
    }{
        //test-caseの作成
        {
            testname:   "正常系",
            content:    "sampleText",
            wantError:  false,
        },
        {
            testname:   "異常系(content無し)",
            content:    "",
            wantError:  true,
        },
    }
    //forでテストケースを回す
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            _, err := user.CreateTodo(tt.content)
            if tt.wantError{
                //tt.wantError=trueの場合
                assert.Error(t, err, "エラーを期待していたが、エラーが帰らない %v")
            }else{
                //tt.wantError=falseの場合
                assert.NoError(t, err, "エラーが発生してしまった %v", err)
            }
        })
    }
}

