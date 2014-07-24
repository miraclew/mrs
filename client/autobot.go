package main

import (
	"github.com/miraclew/mrs/missle/model"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	DSN = "root@tcp(localhost:3306)/mr?charset=utf8"
)

func main() {
	log.Println("I'm Optimus Prime, we're here, we're waiting.")
	db := model.InitDb(DSN)
	// model.SeedData(db)

	var users []model.User
	_, err := db.Select(&users, "select * from users order by id limit 1")
	checkErr(err, "Select failed")
	// log.Println("All rows:")
	for _, u := range users {
		client := DefaultClient{user: u}
		// log.Printf("%#v\n", client)
		go client.run()
	}

	exitChan := make(chan int)
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		exitChan <- 1
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-exitChan
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
