package utils

import(
    "fmt"
    "net/http"
    )

func JsonError(w http.ResponseWriter, status int, msg string){
    //w.HeaderにBody情報にjsonを返すことを宣言
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, msg)))
}