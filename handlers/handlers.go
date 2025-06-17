package handlers

import(
    "net/http"
    )
    
func GetTodos( w http.ResponseWriter, r *http.Request){
    userIDStr := r.URL.Query().Get("userid")
    if userIDStr ==""{
        http.Error(w, "useridが存在しません", http.StatusNotFound)
        return 
        //↑ここでreturnしないと、ResponseWriterに"useridが存在しません"と伝えたにもかかわらず、この後の処理も実行される！
        //必要に応じてstrconv.Atoi
    }
}