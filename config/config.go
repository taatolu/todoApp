package config

import(
    "log"
    "github.com/go-ini/ini"
    "fmt"
    )
    

type ConfigList struct{
    Logfile string
}

var Config ConfigList

func LoadConfig(){
    cfg, err := ini.Load("config.ini")
    
    if err != nil{
        log.Fatalln(err)
    }
    
    Config = ConfigList{
        Logfile:  cfg.Section("web").Key("logfile").String(),
    }
    
    fmt.Println(Config.Logfile)
}