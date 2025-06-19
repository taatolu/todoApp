package models

import(
    "testing"
    "github.com/stretchr/testify/assert"
    "main/config"
    )

func TestInitDB(t *testing.T){
    cfg, err := config.LoadConfig("test")
    err = InitDB(cfg)
    assert.NoError(t, err)
}

//test用のDBイニシャライザー
var DB *sql.DB
func InitTestDB()error{
    cfg, err := config.LoadConfig("test")
    if err != nil{
        return fmt.Errorf("TESTDB接続エラー %w", err)
    }
    
    err = InitDB(cfg)
    if err != nil{
        return fmt.Errorf("TESTDB接続エラー %w", err)
    }
}