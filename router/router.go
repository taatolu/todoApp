package router

import (
    "net/http"
    "main/handlers"
    )

func InitRouters(){
    http.HandleFunc("/api/v1/resource",handlers.GetTodos)
    http.HandleFunc("/api/v2/resource",handlers.GetTodo)
}