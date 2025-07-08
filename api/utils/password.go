package utils

import(
    "golang.org/x/crypto/bcrypt"
    )

//ハッシュ化
func Hash(plaintext string)(string, error){
    hashedPass, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
    return string(hashedPass), err
}

//ハッシュ化した値とのチェック
func HashCheck(plaintext, hash string)bool{
    err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(plaintext))
    return err == nil
}