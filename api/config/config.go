package config

import(
    "os"
    )


type Config struct {
    Logfile  string
    User     string
    Password string
    DBname   string
    Host     string
}


func LoadConfig() *Config {
    return &Config{
        Logfile:  os.Getenv("LOGFILE"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBname:   os.Getenv("DB_NAME"),
        Host:     os.Getenv("DB_HOST"),
    }
}