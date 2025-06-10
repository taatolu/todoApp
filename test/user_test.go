package test

import(
    "testing"
    _ "github.com/stretchr/testify/assert"
    "main/models"
    _ "main/config"
    )

func TestCreateUser(t *testing.T){
    tests := []struct {
        testname    string
        user        *models.User
        wantErr     bool
    }{
        //テストケースの作成
        {
            testname:       "正常系",
            user:           &models.User{
                Name:       "yusaku",
                Email:      "test@example.com",
                Password:   "test123",
            },
            wantErr:        false,
        },
        {
            testname:       "異常系（Nameなし）",
            user:           &models.User{
                Name:       "",
                Email:      "test@example.com",
                Password:   "test123",
            },
            wantErr:        true,
        },
        {
            testname:       "異常系（メアドなし）",
            user:           &models.User{
                Name:       "yusaku",
                Email:      "",
                Password:   "test123",
            },
            wantErr:        true,
        },
        {
            testname:       "異常系（Passなし）",
            user:           &models.User{
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