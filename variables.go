package main

import (
	"github.com/jinzhu/gorm"
	"os"
	"strings"
	"unicode"
)

// IMPORTANT: CHANGE THIS TO FALSE IN DEVELOPMENT
var production = true

// global variable to access the db
var db *gorm.DB

// spotify API variables
var client_id string
var client_secret string
var scopes string

// Postgresql connection variables
var pg_url string
var pg_host string
var pg_database string
var pg_user string
var pg_port string
var pg_password string

// output port for the webserver to connect to
var port string

// initialize global variables - either by getting them from environment variables (production)
// or manually setting them in this function (development)
func initialize() {
	if production{
		client_id = os.Getenv("SPOTIFY_CLIENT_ID")
		client_secret = os.Getenv("SPOTIFY_CLIENT_SECRET")
		scopes = os.Getenv("SPOTIFY_SCOPES")
		pg_url = os.Getenv("DATABASE_URL")
		port = os.Getenv("PORT")
	} else{
		// IMPORTANT: IF IN DEVELOPMENT, CHANGE THESE AND FILL THESE IN YOURSELF
		client_id = ""
		client_secret = ""
		scopes = ""
		pg_url = ""
		port = ""
	}

	// unpack database URL into variables for connection
	// assuming URL will be of the form postgres://<user>:<password>@<host_url>:<port>/<database_name>
	pg_fields := strings.FieldsFunc(pg_url, func(c rune) bool {
		return unicode.IsSpace(c) || c==':' || c=='/' || c=='@'
	})

	pg_user = pg_fields[1]
	pg_password = pg_fields[2]
	pg_host = pg_fields[3]
	pg_port = pg_fields[4]
	pg_database = pg_fields[5]
}