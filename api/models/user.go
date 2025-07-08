package models

import(
    "time"
    "fmt"
    "errors"
    )
    

type User struct {
    ID          int
    UUID        string
    Name        string
    Email       string
    Password    string
    CreatedAt   time.Time
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
    createdat) values ($1,$2,$3,$4,$5) returning id`
    
    err = DB.QueryRow(cmd,
        createUUID(),
        u.Name,
        u.Email,
        utils.Hash(u.Password),
        time.Now()).Scan(&u.ID)
    
    if err != nil {
        return fmt.Errorf("error from CreateUser %w", err)
    }
    return nil
}

func GetUser (userID int)(user *User, err error){
    user = &User{}
    cmd := "select id, uuid, name, email, createdat from users where id = $1"
    err = DB.QueryRow(cmd,userID).Scan(
        &user.ID,
        &user.UUID,
        &user.Name,
        &user.Email,
        &user.CreatedAt)
    if err != nil {
        return nil, fmt.Errorf("error from GetUser: %w", err)
    }

    return user, nil
}

func (u *User) UpdateUser()(err error){
    cmd := "UPDATE users SET name=$2,email=$3 WHERE id=$1"
    //入力項目の不足確認
    ///不足項目のリストアップ
    var missing []string
    if u.ID ==0{
        missing = append(missing, "user.ID")
    }
    if u.Name ==""{
        missing = append(missing, "user.Name")
    }
    if u.Email ==""{
        missing = append(missing, "user.Email")
    }
    //不足項目があった場合は何が不足しているか返す
    if len(missing)>0{
        return fmt.Errorf("入力項目不足 :%v", missing)

    }
    
    _, err = DB.Exec(cmd, u.ID, u.Name, u.Email)
    if err != nil {
        return fmt.Errorf("error from UpdateUser %w", err)
    }
    return nil
}

func DeleteUser(id int)(err error){
    cmd := "DELETE FROM users WHERE id = $1"
    _, err = DB.Exec(cmd,id)
    if err != nil{
        return fmt.Errorf("error from DeleteUser %w", err)
    }
    return nil
}