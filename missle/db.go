package missle

import (
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	GENDER_UNKNOW = 0
	GENDER_MALE   = 1
	GENDER_FEMALE = 2
)

type User struct {
	Uid       int64  `json:"user"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Cellphone string `json:"cellphone"`
	Avatar    string `json:"avatar"`
	Gender    int    `json:"gender"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}

var mc *memcache.Client
var db *sql.DB

func init() {
	mc = memcache.New(MC)
	var err error
	db, err = sql.Open("mysql", DSN)
	if err != nil {
		fmt.Println("........")
		// panic(err.Error())
		// log.Fatal(err.Error())
	} else {
		fmt.Println("+++++++")
	}
}

func FindUserByIdCached(uid int64) (u *User) {
	k := fmt.Sprintf("user:%d", uid)
	it, err := mc.Get(k)
	if err != nil && err != memcache.ErrCacheMiss {
		log.Printf("Error: memcache get failed: %s", err)
		return
	}

	if err == memcache.ErrCacheMiss {
		u = FindUserById(uid)
		v, _ := json.Marshal(u)
		mc.Set(&memcache.Item{Key: k, Value: v})
		return
	}

	u = new(User)
	err = json.Unmarshal(it.Value, u)
	if err != nil {
		log.Println(err)
	}

	return
}

func FindUserById(uid int64) (u *User) {
	user := new(User)
	err := db.QueryRow("select uid,username,password,avatar,cellphone,gender,created from users where uid=?", uid).Scan(
		&user.Uid, &user.UserName, &user.Password, &user.Avatar, &user.Cellphone, &user.Gender, &user.Created)

	if err == sql.ErrNoRows {
		return nil
	}

	return user
}

func FindUserByCredential(username string, password string) (u *User) {
	user := new(User)
	err := db.QueryRow("select uid,username,password,avatar,cellphone,gender,created from users where username=? and password=?", username, password).Scan(
		&user.Uid, &user.UserName, &user.Password, &user.Avatar, &user.Cellphone, &user.Gender, &user.Created)

	if err == sql.ErrNoRows {
		return nil
	}

	return user
}

func GetUidByUserName(username string) (uid int64) {
	uid = 0
	key := fmt.Sprintf("username:%s", username)
	it, err := mc.Get(key)

	if err == memcache.ErrCacheMiss {
		db.QueryRow("select uid from users where username=?", username).Scan(&uid)
		buf := make([]byte, 8)
		binary.PutVarint(buf, uid)
		mc.Set(&memcache.Item{Key: key, Value: buf})
	}

	if err == nil {
		uid, _ = binary.Varint(it.Value)
	}

	return
}

func (user *User) Save() (err error) {
	var result sql.Result
	if user.Uid == 0 { // insert
		result, err = db.Exec("insert into users (username,password) values (?,?)", &user.UserName, &user.Password)
		if err == nil {
			user.Uid, err = result.LastInsertId()
		}
	} else { // update
		_, err = db.Exec("update users set username=?,password=?,avatar=?,cellphone=?,gender=? where uid=?",
			&user.UserName, &user.Password, &user.Avatar, &user.Cellphone, &user.Gender, &user.Uid)
		key := fmt.Sprintf("user:%d", user.Uid)
		mc.Delete(key)
	}

	return
}
