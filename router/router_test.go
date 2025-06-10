package router

import(
    "net/http"
    "net/http/httptest"
    "testing"
    )

func TestGetTodosHandlar(t *testing.T){
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
            if res.StatusCode != tt.wantStatus{
                t.Errorf("got %d, want %d", res.StatusCode, tt.wantStatus)
            }
        })
    }
}