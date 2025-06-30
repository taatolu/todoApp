package models

import(
    "testing"
    "github.com/stretchr/testify/assert"
    "main/config"
    )

func TestInitDB(t *testing.T){
    cfg, err := config.LoadConfig("test")
    assert.NoError(t, err, "LoadConfigでエラー発生")
    err = InitDB(cfg)
    assert.NoError(t, err, "InitDBでエラー発生")
}
