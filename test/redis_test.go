package redis_test

import (
	"log"
	"os"
	"testing"

	"github.com/0x30c4/ghostbin/pkg/env"
	"github.com/0x30c4/ghostbin/pkg/logger"
	"github.com/0x30c4/ghostbin/pkg/redis"
	"github.com/alicebob/miniredis/v2"
)

func TestRedisConnection(t *testing.T) {
  env.EnvInit()

  // changing the defualt logs log file and
  // setting up the logger
  env.EnvInit()
  env.LOG_FILE = "../logs/test.log"

  err := os.Remove(env.LOG_FILE)
  defer os.Remove(env.LOG_FILE)

	logger.LoggerInit(env.LOG_FILE, log.Ldate | log.Ltime | log.Lshortfile)

  // creating mock redis server
  redisTest := miniredis.RunT(t)

  env.REDIS_ADDRESS = redisTest.Addr()

  err = redis.InitRedis()

  if err != nil {
    t.Errorf("Error at redis.InitRedis() = %s; want nil", err)
  }

  // check if invalid server throws any error or not
  env.REDIS_ADDRESS = "invalid_address:1212"

  err = redis.InitRedis()

  if err == nil {
    t.Errorf("redis.InitRedis() didn't throw any errors")
  }

}

func TestGetPasteRDB(t *testing.T) {

  // setting up the env and logger for the
  // mock redis client.
  env.EnvInit()
  env.LOG_FILE = "../logs/test.log"

  err := os.Remove(env.LOG_FILE)
  defer os.Remove(env.LOG_FILE)

	logger.LoggerInit(env.LOG_FILE, log.Ldate | log.Ltime | log.Lshortfile)

  // creating mock redis server
  redisTest := miniredis.RunT(t)

  env.REDIS_ADDRESS = redisTest.Addr()

  err = redis.InitRedis()

  if err != nil {
    t.Errorf("Error at redis.InitRedis() = %s; want nil", err)
  }

  if err != nil {
    t.Errorf("Error at %s", err)
  }

  var pasteModel redis.PasteModel

  pasteModel.PasteId = "test"
  pasteModel.BurnAfter = 10
  pasteModel.ReadCount = 2
  pasteModel.DeepUrl = 10
  pasteModel.Secret = "password"

  err = redis.InitRedis()

  if err != nil {
    t.Errorf("Error at redis.InitRedis() = %s; want nil", err)
  }

  // paste doesn't exist
  err = redis.PutPasteRDB(pasteModel)
  if err != nil {
    t.Errorf("Error at redis.PutPasteRDB(%s) = %v; want nil", pasteModel.PasteId, err)
  }

  // main GetPasteRDB testing starts from here

  // for a known key
  ok, err := redis.GetPasteRDB(pasteModel.PasteId)
  if err != nil {
    t.Errorf(
      "Error: redis.GetPasteRDB(%s) = %t, %s; want true, nil",
      pasteModel.PasteId, ok, err,
      )
  }

  if !ok {
    t.Errorf(
      "Error: redis.GetPasteRDB(%s) = %t, %s; want true, nil",
      pasteModel.PasteId, ok, err,
      )
  }

  // test for an unknown key
  ok, err = redis.GetPasteRDB("key_doesn't exitst")
  if err != nil {
    t.Errorf(
      "Error: redis.GetPasteRDB(%s) = %t, %s; want false, nil",
      pasteModel.PasteId, ok, err,
      )
  }

  if ok {
    t.Errorf(
      "Error: redis.GetPasteRDB(%s) = %t, %s; want true, nil",
      pasteModel.PasteId, ok, err,
      )
  }

  // test for readcount less than 1
  pasteModel.ReadCount = 0

  // put the paste data
  err = redis.PutPasteRDB(pasteModel)
  if err != nil {
    t.Errorf("Error at PutPasteRDB(%s) = %v; want nil", pasteModel.PasteId, err)
  }

  ok, err = redis.GetPasteRDB(pasteModel.PasteId)
  if err != nil {
    t.Errorf(
      "Error: redis.GetPasteRDB(%s) = %t, %s; want true, nil",
      pasteModel.PasteId, ok, err,
      )
  }

  if !ok {
    t.Errorf(
      "Error: redis.GetPasteRDB(%s) = %t, %s; want true, nil",
      pasteModel.PasteId, ok, err,
      )
  }
}


func TestIsPasteExist(t *testing.T) {

  // setting up the env and logger for the
  // mock redis client.
  env.EnvInit()
  env.LOG_FILE = "../logs/test.log"

  err := os.Remove(env.LOG_FILE)
  defer os.Remove(env.LOG_FILE)

	logger.LoggerInit(env.LOG_FILE, log.Ldate | log.Ltime | log.Lshortfile)

  // creating mock redis server
  redisTest := miniredis.RunT(t)

  env.REDIS_ADDRESS = redisTest.Addr()

  err = redis.InitRedis()

  if err != nil {
    t.Errorf("Error at %s", err)
  }

  var pasteModel redis.PasteModel

  pasteModel.PasteId = "test"
  pasteModel.BurnAfter = 10
  pasteModel.ReadCount = 2
  pasteModel.DeepUrl = 10
  pasteModel.Secret = "password"

  err = redis.InitRedis()

  err = redis.PutPasteRDB(pasteModel)
  if err != nil {
    t.Errorf("Error at PutPasteRDB(%s) = %v; want nil", pasteModel.PasteId, err)
  }

  // main IsPasteExist testing starts from here

  ok := redis.IsPasteExist(pasteModel.PasteId)
  if !ok {
    t.Errorf("Error at redis.IsPasteExist(%s) = %v; want true", pasteModel.PasteId, ok)
  }

  ok = redis.IsPasteExist("key_that_doesn't_exists")
  if ok {
    t.Errorf("Error at redis.IsPasteExist(%s) = %v; want false", pasteModel.PasteId, ok)
  }
}


func TestGetNewPasteID(t *testing.T) {

  // setting up the env and logger for the
  // mock redis client.
  env.EnvInit()
  env.LOG_FILE = "../logs/test.log"

  err := os.Remove(env.LOG_FILE)
  defer os.Remove(env.LOG_FILE)

	logger.LoggerInit(env.LOG_FILE, log.Ldate | log.Ltime | log.Lshortfile)

  // creating mock redis server
  redisTest := miniredis.RunT(t)

  env.REDIS_ADDRESS = redisTest.Addr()

  err = redis.InitRedis()

  if err != nil {
    t.Errorf("Error at %s", err)
  }

  err = redis.InitRedis()

  if err != nil {
    t.Errorf("Error at redis.InitRedis() = %s; want nil", err)
  }

  // main GetNewPasteID testing starts from here

  pasteId, err := redis.GetNewPasteID()

  if err != nil {
    t.Errorf("Error at redis.GetNewPasteID() = %s, %s; want pasteId, nil", pasteId, err)
  }

  if len(pasteId) == 0 {
    t.Errorf("Error at redis.GetNewPasteID() = %s, %s; want pasteId, nil", pasteId, err)
  }

}


func TestDeletePasteRDB(t *testing.T) {

  // setting up the env and logger for the
  // mock redis client.
  env.EnvInit()
  env.LOG_FILE = "../logs/test.log"

  err := os.Remove(env.LOG_FILE)
  defer os.Remove(env.LOG_FILE)

	logger.LoggerInit(env.LOG_FILE, log.Ldate | log.Ltime | log.Lshortfile)

  // creating mock redis server
  redisTest := miniredis.RunT(t)

  env.REDIS_ADDRESS = redisTest.Addr()

  err = redis.InitRedis()

  if err != nil {
    t.Errorf("Error at redis.InitRedis() = %s; want nil", err)
  }

  if err != nil {
    t.Errorf("Error at %s", err)
  }

  var pasteModel redis.PasteModel

  pasteModel.PasteId = "test"
  pasteModel.BurnAfter = 10
  pasteModel.ReadCount = 2
  pasteModel.DeepUrl = 10
  pasteModel.Secret = "password"

  err = redis.InitRedis()

  if err != nil {
    t.Errorf("Error at redis.InitRedis() = %s; want nil", err)
  }

  err = redis.PutPasteRDB(pasteModel)
  if err != nil {
    t.Errorf("Error at redis.PutPasteRDB(%s) = %v; want nil", pasteModel.PasteId, err)
  }

  // for a known key
  ok, err := redis.GetPasteRDB(pasteModel.PasteId)
  if err != nil {
    t.Errorf(
      "Error: redis.GetPasteRDB(%s) = %t, %s; want true, nil",
      pasteModel.PasteId, ok, err,
      )
  }

  // main GetPasteRDB testing starts from here

  // delete a known pasteId with proper Secret
  ok, err = redis.DeletePasteRDB(pasteModel.PasteId, pasteModel.Secret)

  if err != nil {
    t.Errorf(
      "Error: redis.DeletePasteRDB(%s, %s) = %t, %s; want true, nil",
      pasteModel.PasteId, pasteModel.Secret, ok, err,
      )
  }

  if !ok {
    t.Errorf(
      "Error: redis.DeletePasteRDB(%s, %s) = %t, %s; want true, nil",
      pasteModel.PasteId, pasteModel.Secret, ok, err,
      )
  }

  // delete a known pasteId with incorrect Secret
  ok, err = redis.DeletePasteRDB(pasteModel.PasteId, "wrong_secret")

  if err == nil {
    t.Errorf(
      "Error: redis.DeletePasteRDB(%s, %s) = %t, %s; want false, redis nil",
      pasteModel.PasteId, pasteModel.Secret, ok, err,
      )
  }

  if ok {
    t.Errorf(
      "Error: redis.DeletePasteRDB(%s, %s) = %t, %s; want false, nil",
      pasteModel.PasteId, pasteModel.Secret, ok, err,
      )
  }

  // delete a known pasteId with no previous Secret value

  pasteModel.Secret = ""
  err = redis.PutPasteRDB(pasteModel)
  if err != nil {
    t.Errorf("Error at redis.PutPasteRDB(%s) = %v; want nil", pasteModel.PasteId, err)
  }

  ok, err = redis.DeletePasteRDB(pasteModel.PasteId, "")

  if err != nil {
    t.Errorf(
      "Error: redis.DeletePasteRDB(%s, %s) = %t, %s; want false, nil",
      pasteModel.PasteId, pasteModel.Secret, ok, err,
      )
  }

  if ok {
    t.Errorf(
      "Error: redis.DeletePasteRDB(%s, %s) = %t, %s; want false, nil",
      pasteModel.PasteId, pasteModel.Secret, ok, err,
      )
  }

}
