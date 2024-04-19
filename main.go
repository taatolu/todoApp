package main

import (
    "main/user"
    "main/utils"
    "main/config"
    
    "fmt"
    "log"

    )

func init () {
    config.LoadConfig()
    utils.LoggingSettings(config.Config.Logfile)
}

func main(){
    user1 := user.NewUser("A123","Yusaku")
    fmt.Println(user1)
    
    log.Println("this is a TEST")
    
}