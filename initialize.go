package main

import (
	"os"
	"strings"
	"unicode"
)

// IMPORTANT: CHANGE THIS TO FALSE IN DEVELOPMENT
var production = true

// initialize global variables - either by getting them from environment variables (production)
// or manually setting them in this function (development)
func initialize() {
	if production{
		client_id = os.Getenv("SPOTIFY_CLIENT_ID")
		client_secret = os.Getenv("SPOTIFY_CLIENT_SECRET")
		scopes = os.Getenv("SPOTIFY_SCOPES")
		pg_url = os.Getenv("DATABASE_URL")
		redis_url = os.Getenv("REDISCLOUD_URL")
		port = os.Getenv("PORT")
		api_url = os.Getenv("API_URL")
		frontend_url = os.Getenv("FRONTEND_URL")
	} else{
		// IMPORTANT: IF IN DEVELOPMENT, CHANGE THESE AND FILL THESE IN YOURSELF
		client_id = ""
		client_secret = ""
		scopes = ""
		pg_url = ""
		redis_url = ""
		port = ""
		api_url = ""
		frontend_url = ""
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

	// unpack redis URL into variables for connection
	// assuming URL will be of the form:  redis://rediscloud:<password>@<endpoint>
	redis_fields := strings.FieldsFunc(redis_url, func(c rune) bool {
		return unicode.IsSpace(c) || c==':' || c=='/' || c=='@'
	})

	redis_endpoint = redis_fields[3] + ":" + redis_fields[4]
	redis_password = redis_fields[2]


	// set react path to serve frontend assets out of
	react_path = "/frontend"
}