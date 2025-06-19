package router

import(
    "net/http"
    "net/http/httptest"
    "io/ioutil"
    "testing"
    "main/models"
    "main/config"
    "fmt"
    )

func TestGetTodosHandlar(t *testing.T){
    
    //test毎に初期化が必要なセット
    cfg, err := config.LoadConfig("test")
    if err != nil{
        fmt.Println(err)
        return
    }
    err = models.InitDB(cfg)
    if err != nil{
        fmt.Println(err)
        return
    }
    
    // ここでハンドラ登録
    InitRouters()
    
    tests := []struct {
        name        string
        url         string
        wantStatus  int
    }{
        //テストケース作成
        {
            name:       "正常系",
            url:        "/api/v1/resource?user_id=1",
            wantStatus: http.StatusOK,
        },
        {
            name:       "user_idなし",
            url:        "/api/v1/resource",
            wantStatus: http.StatusBadRequest,
        },
        {
            name:       "存在しないパス",
            url:        "/api/v1/unknown",
            wantStatus: http.StatusNotFound,
        },
    }
    
    for _, tt := range tests{
        t.Run(tt.name, func(t *testing.T){
            req:= httptest.NewRequest("GET", tt.url, nil)
            rec:= httptest.NewRecorder()
            http.DefaultServeMux.ServeHTTP(rec, req)
            res:=rec.Result()
            defer res.Body.Close()
            
            if res.StatusCode != tt.wantStatus{
                bodyBytes, err := ioutil.ReadAll(res.Body)
                if err != nil{
                    t.Errorf("got %d, want %d, errmsg: %v", res.StatusCode, tt.wantStatus, string(bodyBytes))
                }else{
                    t.Errorf("got %d, want %d, errmsg: %v", res.StatusCode, tt.wantStatus, string(bodyBytes))
                }
                
            }
        })
    }
}