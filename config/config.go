package config

import(
    "log"
    "main/utils"
    "github.com/go-ini/ini"

    )
    

type ConfigList struct{
    Logfile string
}

var Config ConfigList


func init () {
    LoadConfig()
    utils.LoggingSettings(Config.Logfile)
}


func LoadConfig(){
    cfg, err := ini.Load("config.ini")
    
    if err != nil{
        log.Fatalln(err)
    }
    
    Config = ConfigList{
        Logfile:  cfg.Section("web").Key("logfile").String(),
    }
    
}