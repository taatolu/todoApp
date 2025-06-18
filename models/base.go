package models

import(
    "crypto/sha1"
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/google/uuid"
    "fmt"
    "main/config"
    )

var DB *sql.DB

const(
    tableNameUser = "users"
    tableNameTodo = "todos"
    )


func InitDB (conf *config.Config) error {
    user := conf.User
    password :=conf.Password
    dbname := conf.DBname
    connStr :=fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
    DB , err := sql.Open("postgres", connStr)
    
    if err != nil{
        return fmt.Errorf("DB接続エラー %w", err)
    }
    
    cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
        id SERIAL,
        uuid VARCHAR,
        name VARCHAR,
        email VARCHAR,
        password VARCHAR,
        createdat TIMESTAMP)`,tableNameUser)
    
    _ , err = DB.Exec(cmdU)
    if err != nil{
        return fmt.Errorf("createUserTableエラー %w", err)
    }
    
    cmdT:= fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
        id  SERIAL,
        content VARCHAR,
        state   VARCHAR,
        userid  INTEGER,
        createdat   TIMESTAMP,
        updatedat   TIMESTAMP)`,tableNameTodo)
        
    _, err = DB.Exec(cmdT)
    if err != nil {
        return fmt.Errorf("createTodoTableエラー %w", err)
    }
    return nil
}

func createUUID()(uuidobj uuid.UUID) {
    uuidobj, _ =uuid.NewUUID()
    return uuidobj
}


func Encrypt(plaintext string) (cryptext string){
    cryptext = fmt.Sprintf("%x",sha1.Sum([]byte(plaintext)))
    return cryptext
}
