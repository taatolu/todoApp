package test

import(
    "testing"
    "github.com/stretchr/testify/assert"
    "main/config"
    "fmt"
    )

func TestLoadConfig(t *testing.T){
    cfg, err := config.LoadConfig("test")
    assert.NoError(t,err,"config読込エラー %v", err)
    assert.NotEmpty(t, cfg.Logfile,"読込んだ結果が空です")
    fmt.Print(cfg)
}