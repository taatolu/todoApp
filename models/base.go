package models

import(
    "crypto/sha1"
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/google/uuid"
    
    "fmt"
    "log"
    "main/config"
    )


var DB *sql.DB

var err error

const(
    tableNameUser = "users"
    tableNameTodo = "todos"
    )


func InitDB (conf *config.Config) {
    user := conf.User
    password :=conf.Password
    dbname := conf.DBname
    connStr :=fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
    DB , err = sql.Open("postgres", connStr)
    
    if err != nil{
        log.Fatalln(err)
    }
    
    cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
        id SERIAL,
        uuid VARCHAR,
        name VARCHAR,
        email VARCHAR,
        password VARCHAR,
        create_at TIMESTAMP)`,tableNameUser)
    
    _ , err = DB.Exec(cmdU)
    if err != nil{
        log.Fatalln(err)
    }
    
    cmdT:= fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
        id  SERIAL,
        content VARCHAR,
        userid  INTEGER,
        create_at   TIMESTAMP,
        update_at   TIMESTAMP)`,tableNameTodo)
        
    _, err = DB.Exec(cmdT)
    if err != nil {
        log.Fatalln(err)
    }
}

func createUUID()(uuidobj uuid.UUID) {
    uuidobj, _ =uuid.NewUUID()
    return uuidobj
}


func Encrypt(plaintext string) (cryptext string){
    cryptext = fmt.Sprintf("%x",sha1.Sum([]byte(plaintext)))
    return cryptext
}
