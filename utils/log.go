package utils

import (
    "os"
    "log"
    "io"
    "fmt"
    )
    

func LoggingSettings(iniLog string){
    logfile , err := os.OpenFile(iniLog,os.O_RDWR|os.O_CREATE|os.O_APPEND,0766)
    if err != nil {
        log.Fatalln(err)
    }
    
    fmt.Printf("ログファイルは%vやでー",iniLog)
    
    mltiLogfile := io.MultiWriter(logfile,os.Stdout)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetOutput(mltiLogfile)
}