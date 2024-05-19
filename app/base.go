package app

import(
    "crypto/sha1"
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/google/uuid"
    
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
        uuid VARCHAR,
        name VARCHAR,
        email VARCHAR,
        password VARCHAR,
        create_at TIMESTAMP)`,tableNameUser)
    
    _ , err := Db.Exec(cmdU)
    if err != nil{
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
