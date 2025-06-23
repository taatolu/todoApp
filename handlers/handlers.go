package handlers

import(
    "net/http"
    "strconv"
    "encoding/json"
    "main/models"
    )
    
func GetTodos( w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    
    //クエリパラメータからuser_id取得
    userIDStr := r.URL.Query().Get("user_id")
    if userIDStr ==""{
        utils.JsonError(w, http.StatusBadRequest, "user_idが入力されていません")
        return 
        //↑ここでreturnしないと、ResponseWriterに"user_idが存在しません"と伝えたにもかかわらず、この後の処理も実行される！
    }
    
    //取得したuser_idを数値型に変換
    num, err := strconv.Atoi(userIDStr)
    if err != nil{
        utils.JsonError(w, http.StatusBadRequest, "user_idが数値ではありません")
        return
    }
    //userの取得
    user, err := models.GetUser(num)
    if err!= nil{
        utils.JsonError(w, http.StatusNotFound, "userが存在しません")
        return
    }
    
    //todosの取得
    todos, err := user.GetTodos()
    if err != nil{
        utils.JsonError(w, http.StatusNotFound, "todoが登録されていません")
        return
    }
    
    err = json.NewEncoder(w).Encode(&todos)
    if err != nil{
        utils.JsonError(w, http.StatusInternalServerError, "GetTodosのエンコード失敗")
        return
    }
    
}

func GetTodo(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    
    todoidStr := r.URL.Query().Get("todo_id")
    if todoidStr == ""{
        utils.JsonError(w, http.StatusBadRequest, "todo_idが入力されていません")
        return
    }
    
    todoid, err := strconv.Atoi(todoidStr)
    if err!= nil{
        utils.JsonError(w, http.StatusBadRequest, "todo_idが数値ではありません")
        return
    }
    
    //todoの取得
    todo, err := models.GetTodo(todoid)
    if err != nil{
        utils.JsonError(w, http.StatusNotFound, "todoが存在しません")
        return
    }
    
    //jsonにエンコード
    err = json.NewEncoder(w).Encode(&todo)
    if err != nil{
        utils.JsonError(w, http.StatusInternalServerError, "GetTodoのエンコード失敗")
        return
    }
}


