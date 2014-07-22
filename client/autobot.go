package main

import (
	"log"
)

const (
	DSN = "root:abc123@tcp(localhost:3306)/mr?charset=utf8"
)

func main() {
	log.Println("I'm Optimus Prime, we're here, we're waiting.")
	db := initDb()
	// seedData(db)

	var users []User
	_, err := db.Select(&users, "select * from users order by id limit 2")
	checkErr(err, "Select failed")
	log.Println("All rows:")
	for x, u := range users {
		log.Printf("    %d: %v\n", x, u)
		client := &DefaultClient{}
		go client.runAs(&u)
	}

	for {

	}
}
