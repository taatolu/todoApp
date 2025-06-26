package handlers

import(
    "net/http"
    "strconv"
    "encoding/json"
    )

func UsersHandler(w http.ResponseWriter, r http.Request){
    switch r.Method{
        case    http.MethodPost: CreateUser(w, r)
        default:    utils.JsonError(w, http.StatusMethodNotAllowed, "リクエストメソッドが不正です")
    }
}

func UserHandler(w http.ResponseWriter, r http.Request){
    switch r.Method{
        case    http.MethodGet: GetUser(w, r)
        case    http.MethodPut: UpdateUser(w, r)
        case    http.MethodDelete: DeleteUser(w, r)
        default: utils.JsonError(w, http.StatusMethodNotAllowed, "リクエストメソッドが不正です")
    }
}