package models

import(
    "testing"
    "github.com/stretchr/testify/assert"
    "main/config"
    "database/sql"
    )

func TestInitDB(t *testing.T){
    //configのテスト
    cfg := config.LoadConfig()
    assert.NotEmpty(t, cfg, "LoadConfigでデータが取得できない")
    //DBのイニシャライズ
    err := InitDB(cfg)
    assert.NoError(t, err, "InitDBでエラー発生")
    
    //DTが適切に作成できているか確認
    rows, err := DB.Query("select to_regclass('public.users'),to_regclass('public.todos')")
    assert.NoError(t, err, "DTの存在確認でエラー %v", err)
    defer rows.Close()
    
    var userTable,todoTable sql.NullString
    for rows.Next(){
        err = rows.Scan(&userTable, &todoTable)
        assert.NoError(t, err, "psqlのテーブル名取得中にエラー発生 %v", err)
        assert.Equal(t, "users", userTable.String, "userTableが存在しない")
        assert.Equal(t, "todos", todoTable.String, "todoTableが存在しない")
    }
}


//CreateUUIDのテスト
func TestCreateUUID(t *testing.T){
    uuid := createUUID()
    assert.NotNil(t, uuid, "uuidの作成に失敗")
}

//Encryptのテスト
func TestEncrypt(t *testing.T){
    newText := Encrypt("あああ")
    assert.NotEmpty(t, newText, "ハッシュ化されたデータが帰ってない")
    assert.NotEqual(t, "あああ", newText, "ハッシュ化されず、当初の文字のまま")
}


