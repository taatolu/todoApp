package handlers

import(
    "net/http"
    "strconv"
    "encoding/json"
    "main/models"
    )
    
func GetTodos( w http.ResponseWriter, r *http.Request){
    //w.HeaderにBody情報にjsonを返すことを宣言
    w.Header().Set("Content-Type", "application/json")
    
    //クエリパラメータからuser_id取得
    userIDStr := r.URL.Query().Get("user_id")
    if userIDStr ==""{
        http.Error(w, `{"error":"user_idが存在しません"}`, http.StatusBadRequest)
        return 
        //↑ここでreturnしないと、ResponseWriterに"user_idが存在しません"と伝えたにもかかわらず、この後の処理も実行される！
    }
    
    //取得したuser_idを数値型に変換
    num, err := strconv.Atoi(userIDStr)
    if err != nil{
        http.Error(w, `{"error":"user_idの値が不正です"}`, http.StatusBadRequest)
        return
    }
    //userの取得
    user, err := models.GetUser(num)
    if err!= nil{
        http.Error(w,`{"error":"userが存在しません"}`, http.StatusBadRequest)
        return
    }
    
    //todosの取得
    todos, err := user.GetTodos()
    if err != nil{
        http.Error(w, `{"error":"todoが登録されていません"}`, http.StatusBadRequest)
        return
    }
    
    err = json.NewEncoder(w).Encode(&todos)
    if err != nil{
        http.Error(w, `{"error":"エンコード失敗"}`, http.StatusInternalServerError)
        return
    }
    
    
    
}

