package handlers

import(
    "net/http"
    "encoding/json"
    "main/utils"
    "main/models"
    "strconv"
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
    newuser := &models.User{
        Name:   req.Username,
        Email:  req.Email,
        Password:   req.Pass,
    }
    //User作成
    if err = newuser.CreateUser(); err!= nil{
        utils.JsonError(w, http.StatusInternalServerError, "Userの作成に失敗しました")
        return
    }
    
    //作成成功
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message":"userの作成に成功しました"}`))
    return
}

func GetUser(w http.ResponseWriter, r *http.Request){
    //クエリパラメータからuser_idを取得
    userIDStr := r.URL.Query().Get("user_id")
    if userIDStr==""{
        utils.JsonError(w, http.StatusBadRequest, "user_idに値が登録されていません")
        return
    }
    
    //user_idを数値に変換
    userid, err := strconv.Atoi(userIDStr)
    if err != nil{
        utils.JsonError(w, http.StatusBadRequest, "user_idの値が数値ではありません")
        return
    }
    
    //UserをDBから取得
    user, err := models.GetUser(userid)
    if err != nil{
        utils.JsonError(w, http.StatusBadRequest, "user_idのuserが存在しません")
        return
    }
    
    //取得したUser構造体をjsonデータにエンコード&ResponseWriterに返す
    if err = json.NewEncoder(w).Encode(user); err != nil{
        utils.JsonError(w, http.StatusInternalServerError, "user情報のエンコード失敗")
        return
    }
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
    //クエリパラメータからuser_idを取得
    userIDStr:= r.URL.Query().Get("user_id")
    if userIDStr == ""{
        utils.JsonError(w, http.StatusBadRequest, "user_idに値が登録されていません")
        return
    }
    
    //useIDStrを数値に変換
    userid, err := strconv.Atoi(userIDStr)
    if err != nil{
        utils.JsonError(w, http.StatusBadRequest, "user_idが数値ではありません")
        return
    }
    
    //userの取得
    user, err := models.GetUser(userid)
    if err != nil{
        utils.JsonError(w, http.StatusBadRequest, "user_idのuserが存在しません")
        return
    }
    
    //request.BodyのJsonを取得
    ///受け取るための構造体を作成
    var req struct{
        Name    string  `json:"name"`
        Email   string  `json:"email"`
    }
    ///r.Bodyの内容を上記構造体（req）に登録
    if err = json.NewDecoder(r.Body).Decode(&req); err != nil{
        utils.JsonError(w, http.StatusInternalServerError, "Updateしたい内容のパースに失敗")
        return
    }
    //UserのUpdate作業
    ///取得したuserに、r.Bodyから取得した"更新したい値"を代入
    user.Name=req.Name
    user.Email=req.Email
    ///Update処理
    if err = user.UpdateUser(); err != nil{
        utils.JsonError(w, http.StatusBadRequest, "User情報のUpdateに失敗")
        return
    }
    
    //Update成功
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message":"UserのUpdate成功"}`))
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
    //クエリパラメータからuser_idを取得
    userIDStr := r.URL.Query().Get("user_id")
    if userIDStr == ""{
        utils.JsonError(w, http.StatusBadRequest, "user_idに値が登録されていません")
        return
    }
    //userIDStrを数値に変換
    userid, err := strconv.Atoi(userIDStr)
    if err != nil{
        utils.JsonError(w, http.StatusBadRequest, "user_idが数値ではありません")
        return
    }
    //DeleteUserで削除
    if err = models.DeleteUser(userid); err!= nil{
        utils.JsonError(w, http.StatusInternalServerError, "Userを削除できませんでした")
        return
    }
    //削除成功コメント
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "userの削除成功"}`))
    return
}

