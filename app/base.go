package app

import(
    "database/sql"
    _ "github.com/lib/pq"
    
    "fmt"
    "log"
    )


var Db *sql.DB

var err error

const(
    tableNameUser = "users"
    )


func init () {
    Db , err = sql.Open("postgres", "user=ubuntu password=ubuntu dbname=ubuntu sslmode=disable")
    
    if err != nil{
        log.Fatalln(err)
    }
    
    cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
        id SERIAL,
        name VARCHAR)`,tableNameUser)
    
    _ , err := Db.Exec(cmdU)
    if err != nil{
        log.Fatalln(err)
    }
}