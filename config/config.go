package config

import(
    "log"
    "main/utils"
    "github.com/go-ini/ini"

    )
    

type WebConfigList struct{
    Logfile string
}

type DbConfigList struct{
    User string
    Password string
    Dbname string
}

var WebConfig WebConfigList
var DbConfig DbConfigList


func init () {
    LoadConfig()
    utils.LoggingSettings(WebConfig.Logfile)
}


func LoadConfig(){
    cfg, err := ini.Load("config.ini")
    
    if err != nil{
        log.Fatalln(err)
    }
    
    WebConfig = WebConfigList{
        Logfile:  cfg.Section("web").Key("logfile").String(),
    }
    
    DbConfig = DbConfigList{
        User:  cfg.Section("db").Key("user").String(),
        Password: cfg.Section("db").Key("password").String(),
        Dbname: cfg.Section("db").Key("dbname").String(),
    }
}