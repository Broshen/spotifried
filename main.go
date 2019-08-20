package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gorilla/mux"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Aw fuck shit an error occured. Fuck me in the ass goddamnit. Can't even get one thing right today...")
}

// reactScriptsHandlerSystem is an http.FileSystem that serves
// a react build static files, and lets react-router serve any routes
// not found
type reactScriptsHandlerSystem struct {
	http.FileSystem
}

// Open is a wrapper around the Open method of the embedded FileSystem
// that serves a 403 permission error when name has a file or directory
// with whose name starts with a period in its path.
func (fs reactScriptsHandlerSystem) Open(name string) (http.File, error) {
	fmt.Println("open", name)
	file, err := fs.FileSystem.Open(name)

	// if a file is not found, assume it is a react-router route, and serve the index page
	if err != nil {
		fmt.Println("oh no ", name, " created error ", err)
		index, err := fs.FileSystem.Open("/index.html")
		if err != nil {
			fmt.Println("oh no serving index created error ", err)
			return nil, err
		}
		return index, nil
	}
	return file, nil
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

	r.PathPrefix("/").Handler(http.FileServer(reactScriptsHandlerSystem{http.Dir("./frontend/build/")}))
	http.Handle("/", r)

	fmt.Println("Online")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}