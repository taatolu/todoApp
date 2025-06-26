package handlers

import(
    "net/http"
    "strconv"
    "encoding/json"
    "main/utils"
    )

func UsersHandler(w http.ResponseWriter, r *http.Request){
    switch r.Method{
        case    http.MethodPost: CreateUser(w, r)
        default:    utils.JsonError(w, http.StatusMethodNotAllowed, "リクエストメソッドが不正です")
    }
}

func UserHandler(w http.ResponseWriter, r *http.Request){
    switch r.Method{
        case    http.MethodGet: GetUser(w, r)
        case    http.MethodPut: UpdateUser(w, r)
        case    http.MethodDelete: DeleteUser(w, r)
        default: utils.JsonError(w, http.StatusMethodNotAllowed, "リクエストメソッドが不正です")
    }
}

func CreateUser(w http.ResponseWriter, r *http.Request){
    //request.Bodyから受取る構造体を作成
    var req struct{
        Username  string  `json:"username"`
        Email   string  `json:"email"`
        Pass    string  `json:"password"`
    }
    
    //requestのBodyを読み取り
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil{
        utils.JsonError(w, http.StatusInternalServerError, "リクエストボディのパースに失敗")
        return
    }
    //入力項目の不足確認
    ///不足項目のリストアップ
    var missing []string        //不足項目を保存すすリストを作成
    if req.Username==""{
        missing = append(missing, "username")
    }
    if req.Email==""{
        missing = append(missing, "email")
    }
    if req.Pass==""{
        missing = append(missing, "password")
    }
    ///不足項目があった場合は何が不足しているか返す
    if len(missing)>0{
        resp := map[string]interface{}{
            "error": "必要な情報が不足しています",
            "missing_fields": missing,
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        if err = json.NewEncoder(w).Encode(resp); err!=nil{
            utils.JsonError(w, http.StatusInternalServerError, "項目不足のエラー作成時にエンコード失敗")
        }
        return
    }
    
    //取得したBodyの情報を構造体にマッピング
    newuser := &User{
        Name:   req.Usernaem,
        Email:  rreq.Email,
        Password:   req.Pass
    }
    //User作成
    if err = newuser.CreateUser(); err!= nil{
        utils.JsonError(w, http.StatusInternalServerError, "Userの作成に失敗しました")
    }
    
    //作成成功
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader("http.StatusOK")
    w.Write([]byte(`{"message": "userの作成に成功しました"}`))
    
}

