package main

import (
    "main/user"
    "main/config"
    "fmt"
    )

func main(){
    user1 := user.NewUser("A123","Yusaku")
    fmt.Println(user1)
    
    config.LoadConfig()
}