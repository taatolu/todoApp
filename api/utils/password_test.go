package utils

import(
    "testing"
    "github.com/stretchr/testify/assert"
    )

//ハッシュ化のテスト
func TestHash(t *testing.T){
    password := "samplePassword"
    hashedPass, err := Hash(password)
    assert.NoError(t, err, "ハッシュ化失敗 %v", err)
    assert.NotEqual(t, password, hashedPass, "ハッシュ化されたpasseordと渡した値が同じ")
}

//ハッシュチェックのテスト
func TestHashCheck(t *testing.T){
    password := "samplePassword"
    hashedPass, err := Hash(password)
    //正しいパスワードの場合
    if !HashCheck(password, hashedPass){
        t.Errorf("正しいパスワードを渡したのに、認証失敗")
    }
    //不正なパスワードの場合
    if HashCheck("tekito-", hashedPass){
        t.Errorf("不正なパスワードを渡したのに、認証成功")
    }
}