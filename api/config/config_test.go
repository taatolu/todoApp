package config

import(
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
    )

func TestLoadConfig(t *testing.T){
    cfg := LoadConfig()
    assert.NotEmpty(t, cfg.Logfile,"読込んだ結果が空です")
    fmt.Print(cfg)
}