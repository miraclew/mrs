package model

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

const (
	GENDER_UNKNOW = 0
	GENDER_MALE   = 1
	GENDER_FEMALE = 2
)

type User struct {
	Id        int64
	UserName  string
	NickName  string
	Password  string
	Cellphone string
	Avatar    string
	Gender    int
	LastLogin int64
	CreatedAt int64
	UpdatedAT int64
}

func InitDb(dsn string) *gorp.DbMap {
	db, err := sql.Open("mysql", dsn)
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func seedData(db *gorp.DbMap) {
	names := []string{"Optimus Prime", "BumbleBee", "Hound", "Drift", "Crosshair", "Grimilock", "Slug", "Strafe", "Scron"}
	avatar := []string{
		"http://a.hiphotos.baidu.com/baike/c0%3Dbaike116%2C5%2C5%2C116%2C38/sign=4dc8e79100087bf469e15fbb93ba3c49/08f790529822720e5cae3a2079cb0a46f31fab8c.jpg",
		"http://baike.baidu.com/picture/10900924/11204216/0/b2de9c82d158ccbfd2f169e01bd8bc3eb035419f?fr=lemma&ct=single",
		"http://baike.baidu.com/picture/10900924/11204216/0/b2de9c82d158ccbfd6436de01bd8bc3eb0354169?fr=lemma&ct=single",
		"http://baike.baidu.com/picture/10900924/11204216/0/a8ec8a13632762d0cbc87f65a2ec08fa503dc672?fr=lemma&ct=single",
		"http://baike.baidu.com/picture/10900924/11204216/0/5fdf8db1cb134954a5212fb5544e9258d1094a20?fr=lemma&ct=single",
		"http://baike.baidu.com/picture/10900924/11204216/0/35a85edf8db1cb13dfa00ec8df54564e93584b1b?fr=lemma&ct=single",
		"http://baike.baidu.com/picture/10900924/11204216/0/e850352ac65c103806d0f119b0119313b17e89ba?fr=lemma&ct=single",
		"http://baike.baidu.com/picture/10900924/11204216/0/b3119313b07eca8031001d3c932397dda04483e7?fr=lemma&ct=single",
		"http://baike.baidu.com/picture/10900924/11204216/0/e850352ac65c10381d6aea19b0119313b17e89d0?fr=lemma&ct=single",
	}
	for i := 0; i < len(names); i++ {
		user := &User{
			UserName: names[i],
			NickName: names[i], Avatar: avatar[i],
			Password:  "bot",
			CreatedAt: time.Now().UnixNano(),
		}
		db.Insert(user)
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
