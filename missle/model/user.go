package model

import (
// "database/sql"
// "encoding/json"
// "fmt"
// _ "github.com/go-sql-driver/mysql"
// "github.com/miraclew/mrs/util"
// "log"
)

const (
	GENDER_UNKNOW = 0
	GENDER_MALE   = 1
	GENDER_FEMALE = 2
)

type UserModel struct {
	Uid       int64  `json:"user"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Cellphone string `json:"cellphone"`
	Avatar    string `json:"avatar"`
	Gender    int    `json:"gender"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}

// var db *sql.DB

func init() {
	// var err error
	// db, err = sql.Open("mysql", DSN)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
}
