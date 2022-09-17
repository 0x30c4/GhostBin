package env

import (
	"os"
	"strconv"
)

// Loading env variables.

var HOST = os.Getenv("HOST")
var PORT = os.Getenv("PORT")
var DOMAIN = os.Getenv("DOMAIN")
var DATA_DIR = os.Getenv("DATA_DIR")
var PASTE_DATA_PATH = os.Getenv("PASTE_DATA_PATH")

var DB_HOST = os.Getenv("DB_HOST")
var DB_PORT = os.Getenv("DB_PORT")
var DB_NAME = os.Getenv("DB_NAME")

var LOG_FILE = os.Getenv("LOG_FILE")
var RELEASE_MODE, _ = strconv.ParseBool(os.Getenv("RELEASE_MODE"))
