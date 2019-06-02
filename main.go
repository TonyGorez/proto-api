package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

/*
	*** TODO:
	- Add orm
	- add gorilla/mux
	- add user handler
*/

func main() {

	db, err := InitDB()
	if err != nil {
		log.Fatal("Form InitDB", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal("Form closing DB", err)
		}
	}()

	// these following lines are temporary
	var id int
	var email string
	var firstName string
	var lastName string
	queryError := db.QueryRow("SELECT * FROM users").Scan(&id, &firstName, &lastName, &email)
	if queryError != nil {
		panic(queryError)
	}
	fmt.Println(id, " | ", email, " | ", firstName, " | ", lastName)


}
