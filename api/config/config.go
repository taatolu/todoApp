package config

import(
    "fmt"
    "github.com/go-ini/ini"
    "runtime"
    "path/filepath"
    )


type Config struct{
    Logfile string
    User string
    Password string
    DBname string
}

func getProjectRoot()string{
    _, filename, _, _ := runtime.Caller(0)
    return filepath.Dir(filepath.Dir(filename))
}


func LoadConfig(section string)(*Config, error){
    root:= getProjectRoot()
    cfgpath:= filepath.Join(root,"config.ini")
    cfg, err := ini.Load(cfgpath)
    
    if err != nil{
        return nil, fmt.Errorf("iniファイル読込エラー：%w", err)
    }
    
    return &Config {
        Logfile:    cfg.Section(section).Key("logfile").String(),
        User:  cfg.Section(section).Key("user").String(),
        Password: cfg.Section(section).Key("password").String(),
        DBname: cfg.Section(section).Key("dbname").String(),
    }, nil
}