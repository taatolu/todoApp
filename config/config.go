package config

import(
    "fmt"
    _ "main/utils"
    "github.com/go-ini/ini"

    )


type Config struct{
    Logfile string
    User string
    Password string
    DBname string
}


func init () {
    //utils.LoggingSettings(WebConfig.Logfile)
}


func LoadConfig(section string)(*Config, error){
    cfg, err := ini.Load("config.ini")
    
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