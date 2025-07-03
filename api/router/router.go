package router

import (
    "net/http"
    "main/handlers"
    )

func InitRouters() http.Handler {
    mux := http.NewServeMux()       //新しくマルチプレクサを作成
    mux.HandleFunc("/api/v1/todos", handlers.TodosHandler) //GET,POSTメソッド
    mux.HandleFunc("/api/v1/todos/", handlers.TodoHandler) //GET,PUSH,DELETEメソッド
    mux.HandleFunc("/api/v1/users", handlers.UsersHandler) //GET,POSTメソッド
    mux.HandleFunc("/api/v1/users/", handlers.UserHandler) //GET,UPDATE.DELETEメソッド
    return mux
}