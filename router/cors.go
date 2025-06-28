package router

import("net/http")

func CORSMiddleware(next http.Handler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        // どのオリジンからのアクセスを許可するか
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        // どのHTTPメソッドを許可するか
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        // どんなヘッダーを許可するか
        w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        
        // プリフライトリクエスト（OPTIONS）への特別な対応
        if r.Method=="OPTIONS"{
            w.WriteHeader(http.StatusOK)
            return
        }

        // 通常のリクエストは次のハンドラへ
        next.ServeHTTP(w, r)
    })
}


