package main

import (
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis"
)

// global variable to access the db
var db *gorm.DB

// global variable to access redis
var redis_db *redis.Client

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

// Redis connection variables
var redis_url string
var redis_endpoint string
var redis_password string

// React folder path
var react_path string

// Frontend/backend URLs
var api_url string
var frontend_url string
