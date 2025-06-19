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
