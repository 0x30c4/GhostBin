package env

import (
	"os"
	"strconv"
)

var HOST string
var PORT string
var DOMAIN string
var LOG_FILE string
var PASTE_DIR string
var REDIS_ADDRESS string
var RELEASE_MODE bool

func EnvInit() {
  HOST = os.Getenv("HOST")
  PORT = os.Getenv("PORT")
  DOMAIN = os.Getenv("DOMAIN")
  LOG_FILE = os.Getenv("LOG_FILE")
  PASTE_DIR = os.Getenv("PASTE_DIR")
  REDIS_ADDRESS = os.Getenv("REDIS_ADDRESS")
  RELEASE_MODE, _ = strconv.ParseBool(os.Getenv("RELEASE_MODE"))
}
