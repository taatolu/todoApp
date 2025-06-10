package test

import(
    "testing"
    "github.com/stretchr/testify/assert"
    "main/models"
    "main/config"
    )

func TestInitDB(t *testing.T){
    cfg, err := config.LoadConfig("test")
    err = models.InitDB(cfg)
    assert.NoError(t, err)
}