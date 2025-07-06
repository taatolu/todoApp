package handlers

import(
    "net/http"
    "strconv"
    "encoding/json"
    "main/models"
    "main/utils"
    )
    
//メソッドの対応をr.Methodで判定
///TodosHandler
func TodosHandler(w http.ResponseWriter, r *http.Request){
    switch r.Method{
        case http.MethodGet: GetTodos(w, r)
        case http.MethodPost: CreateTodo(w, r)
        default: utils.JsonError(w, http.StatusMethodNotAllowed, "リクエストメソッドが不正です")
    }
}

///TodoHandler
func TodoHandler(w http.ResponseWriter, r *http.Request){
    switch r.Method{
        case http.MethodGet: GetTodo(w, r)
        case http.MethodPut: UpdateTodo(w, r)
        default: utils.JsonError(w, http.StatusMethodNotAllowed, "リクエストメソッドが不正です")
    }
}

//GetTodos
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

///CreateTodo
func CreateTodo(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    useridStr := r.URL.Query().Get("user_id")
    if useridStr == ""{
        utils.JsonError(w, http.StatusBadRequest, "user_idが入力されていません")
        return
    }
    userid, err := strconv.Atoi(useridStr)
    if err != nil{
        utils.JsonError(w, http.StatusBadRequest, "user_idが数値ではありません")
        return
    }
    user, err := models.GetUser(userid)
    if err != nil{
        utils.JsonError(w, http.StatusNotFound, "userが存在しません")
        return
    }
    
    //リクエストボディからcontentを取得
    var req struct{
        Content string  `json:"content"`
    }
    
    err = json.NewDecoder(r.Body).Decode(&req)
    if err != nil{
        utils.JsonError(w, http.StatusBadRequest, "リクエストボディのパースに失敗しました")
        return
    }
    if req.Content ==""{
        utils.JsonError(w, http.StatusBadRequest, "contentが入力されていません")
        return
    }
    
    //todoの作成
    _, err = user.CreateTodo(req.Content)
    if err != nil{
        utils.JsonError(w, http.StatusInternalServerError, "todoの作成に失敗")
        return
    }
    //todoの作成成功レスポンス
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "todoの作成に成功しました"}`))
}

//GetTodo
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

//UpdateTodo
func UpdateTodo(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    //クエリパラメータからtodo_id取得
    todoidStr := r.URL.Query().Get("todo_id")
    if todoidStr == ""{
        utils.JsonError(w, http.StatusBadRequest, "todo_idが入力されていません")
        return
    }
    
    todoid, err := strconv.Atoi(todoidStr)
    if err != nil{
        utils.JsonError(w, http.StatusBadRequest, "todo_idが数値ではありません")
        return
    }
    
    todo, err := models.GetTodo(todoid)
    if err != nil{
        utils.JsonError(w, http.StatusNotFound, "todo_idに一致するtodoが存在しません")
        return
    }
    
    //todoのupdate作業
    ///request.Bodyから変更したいcontentを取得
    var req struct{
        Contents    string  `"json":"contents"`
    }
    if err = json.NewDecoder(r.Body).Decode(&req); err!=nil{
        utils.JsonError(w, http.StatusInternalServerError, "変更したい内容のパースに失敗")
        return
    }
    
    if err = todo.UpdateTodo(req.Contents); err!=nil{
        utils.JsonError(w, http.StatusInternalServerError, "todoのUpdateに失敗")
        return
    }
    
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"massage":"todoのUpdateに成功しました"}`))
}
