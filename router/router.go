package router

import (
    "net/http"
    "main/handlers"
    )

func InitRouters(){
    http.HandleFunc("/api/v1/todos", handlers.TodosHandler) //GET,POSTメソッド
    http.HandleFunc("/api/v1/todos/", handlers.TodoHandler) //GET,PUSH,DELETEメソッド
    http.HandleFunc("/api/v1/users", handlers.UsersHandler) //GET,POSTメソッド
    http.HandleFunc("/api/v1/users/", handlers.UserHandler) //GET,UPDATE.DELETEメソッド
}