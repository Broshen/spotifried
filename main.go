package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{}
	fmt.Println("request", r)
	tmpl := template.Must(template.ParseFiles("./frontend/build/index.html"))
	tmpl.Execute(w, response)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
		LastRefreshed: "",
		Songs: "",
		Artists: "",
		Genres: "",

	}
	db.AutoMigrate(&sampleUser)

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/error", errorHandler)
	api.HandleFunc("/authenticate", authHandler)
	api.HandleFunc("/authenticated", authenticatedHandler)
	api.HandleFunc("/fetch/{user_id}", fetchHandler)
	api.HandleFunc("/profile/{user_id}", analyzeHandler)
	api.HandleFunc("/share/{user_id}", shareHandler)
	api.HandleFunc("/compare/{user1_id}/{user2_id}", compareHandler)

	r.PathPrefix("/static").Handler(http.FileServer(http.Dir("./frontend/build/static")))
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("oh no not found", r.URL.String())
	    http.ServeFile(w, r, "./frontend/build/index.html")
	})

	http.Handle("/", r)

	fmt.Println("Online")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}