package redis_test

import (
  "os"
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/0x30c4/ghostbin/pkg/env"
	"github.com/0x30c4/ghostbin/pkg/logger"
)

func TestLogger(t *testing.T) {

  env.EnvInit()
  env.LOG_FILE = "../logs/test.log"

  err := os.Remove(env.LOG_FILE)
  defer os.Remove(env.LOG_FILE)

	logger.LoggerInit(env.LOG_FILE, log.Ldate | log.Ltime | log.Lshortfile)

  dummyLogdata := "Wriging data to log file"
  log.Println(dummyLogdata)

  // read the whole file at once
  b, err := ioutil.ReadFile(env.LOG_FILE)
  if err != nil {
    t.Errorf("Error : %s", err)
  }

  s := string(b)
  // //check whether s contains substring text
  exists := strings.Contains(s, dummyLogdata)

  if !exists {
    t.Errorf(
      "The string '%s' does not exist in the log file '%s.\n",
      dummyLogdata, env.LOG_FILE)
  }

}
