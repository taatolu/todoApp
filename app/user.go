package app

import(
    "time"
    "log"
    )
    

type User struct {
    ID          int
    UUID        string
    Name        string
    Email       string
    Password    string
    CreateAt   time.Time
}

func (u *User) CreateUser()(err error){
    cmd := `insert into users (
    uuid,
    name,
    email,
    password,
    create_at) values ($1,$2,$3,$4,$5)`
    
    _, err = Db.Exec(cmd,
        createUUID(),
        u.Name,
        u.Email,
        Encrypt(u.Password),
        time.Now())
    
    if err != nil {
        log.Fatalln(err)
    }
    return err
}

func GetUser (userID int)(user User, err error){
    cmd := "select * from users where id = $1"
    err = Db.QueryRow(cmd,userID).Scan(
        &user.ID,
        &user.UUID,
        &user.Name,
        &user.Email,
        &user.Password,
        &user.CreateAt)

    return user, err
}