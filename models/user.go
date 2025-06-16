package models

import(
    "time"
    "log"
    "errors"
    )
    

type User struct {
    ID          int
    UUID        string
    Name        string
    Email       string
    Password    string
    CreateAt   time.Time
}

//create user
func (u *User) CreateUser()(err error){
    if u.Name ==""{
        return errors.New("name is required")
    }
    if u.Email ==""{
        return errors.New("email is required")
    }
    if u.Password ==""{
        return errors.New("password is required")
    }
    cmd := `insert into users (
    uuid,
    name,
    email,
    password,
    create_at) values ($1,$2,$3,$4,$5)`
    
    _, err = DB.Exec(cmd,
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

func GetUser (userID int)(user *User, err error){
    user = &User{}
    cmd := "select * from users where id = $1"
    err = DB.QueryRow(cmd,userID).Scan(
        &user.ID,
        &user.UUID,
        &user.Name,
        &user.Email,
        &user.Password,
        &user.CreateAt)

    return user, err
}

func (u *User) UpdateUser()(err error){
    cmd := "UPDATE users SET name=$2,email=$3 WHERE id=$1"
    
    _, err = DB.Exec(cmd, u.ID, u.Name, u.Email)
    if err != nil {
        log.Fatalln(err)
    }
    return err
}

func DeleteUser(id int)(err error){
    cmd := "DELETE FROM users WHERE id = $1"
    _, err = DB.Exec(cmd,id)
    if err != nil{
        log.Fatalln(err)
    }
    return err
}