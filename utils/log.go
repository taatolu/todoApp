package utils

import (
    "os"
    "log"
    "io"
    
    "main/config"
    )
    

func LoggingSettingth(ini-Log string){
    logfile , err := os.OpenFile(ini-Log,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
    if err != nil {
        log.Fatalln(err)
    }
    
    mltiLogfile := io.MultiWriter(logfile,os.Stdout)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetOutput(mltiLogfile)
}