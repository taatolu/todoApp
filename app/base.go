package app

import(
    "database/sql"
    _ "github.com/lib/pq"
    
    "fmt"
    "log"
    "main/config"
    )


var Db *sql.DB

var err error

const(
    tableNameUser = "users"
    )


func init () {
    user := config.DbConfig.User
    password := config.DbConfig.Password
    dbname := config.DbConfig.Dbname
    connStr :=fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
    Db , err = sql.Open("postgres", connStr)
    
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