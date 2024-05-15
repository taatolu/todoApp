package main

import (
    "main/user"
    _ "main/app"
    
    "fmt"
    "log"

    )



func main(){
    user1 := user.NewUser("A123","Yusaku")
    fmt.Println(user1)
    
    log.Println("this is a TEST")
    
}