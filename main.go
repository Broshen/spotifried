package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aw fuck shit an error occured. Fuck me in the ass goddamnit. Can't even get one thing right today...")
}

func main() {

	initialize() // populate GLOBAL variables - see variables.go

	var err error
	db, err = gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s",
			pg_host,
			pg_port,
			pg_user,
			pg_database,
			pg_password,
		),
	)
	defer db.Close()

	if err != nil{
		panic(fmt.Sprintf("Error connecting to database: %s", err))
	}

	sampleUser := User{
		ID: "0",
		DisplayName: "sample",
		AccessToken: "sample",
		RefreshToken: "sample",
	}
	db.AutoMigrate(&sampleUser)

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/error", errorHandler)
	r.HandleFunc("/authenticate", authHandler)
	r.HandleFunc("/authenticated", authenticatedHandler)
	r.HandleFunc("/share/{user_id}", shareHandler)
	r.HandleFunc("/compare/{user1_id}/{user2_id}", compareHandler)
	http.Handle("/", r)

	fmt.Println("Online")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}