package models

import(
    "testing"
    "github.com/stretchr/testify/assert"
    )

func TestCreateUser(t *testing.T){
    tests := []struct {
        testname    string
        user        *User
        wantErr     bool
    }{
        //テストケースの作成
        {
            testname:       "正常系",
            user:           &User{
                Name:       "yusaku",
                Email:      "test@example.com",
                Password:   "test123",
            },
            wantErr:        false,
        },
        {
            testname:       "異常系（Nameなし）",
            user:           &User{
                Name:       "",
                Email:      "test@example.com",
                Password:   "test123",
            },
            wantErr:        true,
        },
        {
            testname:       "異常系（メアドなし）",
            user:           &User{
                Name:       "yusaku",
                Email:      "",
                Password:   "test123",
            },
            wantErr:        true,
        },
        {
            testname:       "異常系（Passなし）",
            user:           &User{
                Name:       "yusaku",
                Email:      "test@example.com",
                Password:   "",
            },
            wantErr:        true,
        },
        
    }
    //テストケースをループで回す
    for _, tt := range tests {
        t.Run(tt.testname, func(t *testing.T){
            err := tt.user.CreateUser()
            if (err != nil) != tt.wantErr {
                t.Errorf("検知したエラー %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestGetUser(t *testing.T){
    //Test用のUserを作成しておく（毎回このUesrをGetする）
    userSample := User{Name:"Get用User", Email:"Get@exam.com", Password:"getexam"}
    err := userSample.CreateUser()
    if err != nil{  t.Errorf("TestGetUser用にUserを作成するところでエラー %w", err)    }

    tests := []struct{
        testname    string
        userid      int
        wantuser    *User
        wantErr     bool
    }{
        //testケースの作成
        {
            //正常系
            testname:   "正常系",
            userid:     userSample.ID,
            wantuser:   &User{
                Name:   userSample.Name,
                Email:  userSample.Email,
            },
            wantErr:    false,
        },
        {
            //異常系
            testname:   "異常系(作成されていないuserid)",
            userid:     userSample.ID +1,
            wantuser:   nil,
            wantErr:    true,
        },
    }
    
    //テストケースをループで回す
    for _, tt := range tests{
        t.Run(tt.testname, func(t *testing.T){
            user, err := GetUser(tt.userid)
            if tt.wantErr{
                //wantErr=true エラーがあった方がよい場合（異常系の場合）
                assert.Error(t, err, "エラーが帰るべきなのにエラーが帰っていない")
                assert.Nil(t, user, "userがnilであるべきなのにnilでない user=%v", user)
            }else{
                //wantErr=false エラーが無いほうが良い場合（正常系の場合）
                assert.NoError(t, err, "エラーが無いほうが良いのに、エラーが発生している")
                assert.NotNil(t, user, "userがnilになっている")
                assert.Equal(t, tt.wantuser.Name, user.Name)
                assert.Equal(t, tt.wantuser.Email, user.Email)
            }

        })
    }
}

func TestUpdateUser(t *testing.T){
    //テーブル駆動テスト用に構造体を作成
    tests := []struct{
        testname    string
        user        *User
        wantErr     bool
    }{
        //テストケース作成
        {
            testname:   "正常系",
            user:       &User{
                ID:     1,
                Name:   "変更後",
                Email:  "changed@example.com",
            },
            wantErr:    false,
        },
        {
            testname:   "異常系(user.IDなし)",
            user:       &User{
                ID:     0,  // int型のデフォルト値
                Name:   "変更後",
                Email:  "changed@example.com",
            },
            wantErr:    true,
        },
        {
            testname:   "異常系(user.Nameなし)",
            user:       &User{
                ID:     1,
                Name:   "",
                Email:  "changed@example.com",
            },
            wantErr:    true,
        },
        {
            testname:   "異常系(user.Emailなし)",
            user:       &User{
                ID:     1,
                Name:   "変更後",
                Email:  "",
            },
            wantErr:    true,
        },
    }
    //テストケースをループで回す
    for _, tt := range tests{
        t.Run(tt.testname, func(t *testing.T){
            err := tt.user.UpdateUser()
            if tt.wantErr{
                //wantErr=true エラーがあった方がよい場合（異常系の場合）
                assert.Error(t, err, "エラーを期待しているのにエラーが帰らない")
            }else{
                //wantErr=false エラーがない方がよい場合（正常系の場合）
                assert.NoError(t, err, "期待していなエラーが発生 %v", err)
            }
        })
    }
}

func TestDeleteUser(t *testing.T){
    tests := []struct{
        testname    string
        userid      int
    }{
        //testケースの作成
        {
            //正常系
            testname:   "正常系",
            userid:     1,
        },
        {
            //異常系
            testname:   "存在しないIDの削除",
            userid:     1000,
        },
    }
    //テストケースをループで回す
    for _, tt := range tests{
        tt := tt // クロージャ対策
        t.Run(tt.testname, func(t *testing.T){
            err := DeleteUser(tt.userid)
            assert.NoError(t, err, "期待していなエラーが発生 %v", err)
            user, err := GetUser(tt.userid)
            assert.Error(t, err, "Userの削除に成功したのでGetすると本来エラーになるはず")
            assert.Nil(t, user, "Userの削除に成功したのでGetすると本来nilになるはず")
        })
    }
}