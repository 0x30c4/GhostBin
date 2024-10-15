package redis

import (
	"context"
	"log"
	"strconv"

	"github.com/0x30c4/ghostbin/pkg/env"
	"github.com/redis/go-redis/v9"
)

type PasteModel struct {
  PasteId   string `redis:"PasteId"`
  BurnAfter uint64 `redis:"BurnAfter"`
  ReadCount uint64 `redis:"ReadCount"`
  DeepUrl   uint8  `redis:"DeepUrl"`
  Secret    string `redis:"Secret"`
}

var REDIS_CLIENT *redis.Client

var ctx = context.Background()

func InitRedis() error {
  // Connect to Redis server
  REDIS_CLIENT = redis.NewClient(&redis.Options{
      Addr:     env.REDIS_ADDRESS, // Redis server address
      Password: "",               // No password set
      DB:       0,                // Use default DB
  })

  // Ping the Redis server to check if connection was successful
  _, err := REDIS_CLIENT.Ping(ctx).Result()

  if err != nil {
    log.Println(err)
    return err
  }
  return nil
}

func GetNewPasteID() (string, error) {

  tx := REDIS_CLIENT.TxPipeline()

  tx.Incr(ctx, "totalPaste")
  totalPaste := tx.Get(ctx, "totalPaste")

  _, err := tx.Exec(ctx)
  if err != nil {
    log.Println(err)
    return "", err
  }

  totalPasteValue, err := totalPaste.Int64()

  if err != nil {
    return "", err
  }

  newPasteID := strconv.FormatInt(int64(totalPasteValue), 16)

  return newPasteID, nil
}

func PutPasteRDB(pasteModel PasteModel) error {

  if pasteModel.ReadCount == 0 {
    return nil
  }

  key := "pasteId:" + pasteModel.PasteId

  err := REDIS_CLIENT.HSet(ctx, key, pasteModel).Err()

  if err != nil {
    return err
  }

  err = REDIS_CLIENT.Do(ctx, "EXPIRE", key, pasteModel.BurnAfter).Err()

  return err
}

func IsPasteExist(pasteId string) bool {
  key := "pasteId:" + pasteId
  exists, err := REDIS_CLIENT.Exists(ctx, key).Result()
  if err != nil {
      log.Println("Error checking key existence:", err)
      return false
  }

  if exists > 0 {
    return true
  } else {
    return false
  }
}


func GetPasteRDB(pasteId string) (bool, error) {
  key := "pasteId:" + pasteId

  res := REDIS_CLIENT.HGetAll(ctx, key)

  err := res.Err()
  if err != nil {
		return false, err
	}

  var pasteModel PasteModel

  err = res.Scan(&pasteModel)

  if (PasteModel{}) == pasteModel {
    return false, nil
  }

	if err != nil {
    return false, err
	}

  if pasteModel.ReadCount <= 1 {
    _ = REDIS_CLIENT.Del(ctx, key)
    return true, nil
  }
  err = REDIS_CLIENT.HIncrBy(ctx, key, "ReadCount", -1).Err()

  if err != nil {
    return false, nil
  }

  return true, nil
}


func DeletePasteRDB(pasteId string, secret string) (bool, error) {
  key := "pasteId:" + pasteId

  secretDB, err := REDIS_CLIENT.HGet(ctx, key, "Secret").Result()

  if err != nil {
    return false, err
  }

  if len(secretDB) == 0 {
    return false, nil
  }

  if secretDB != secret {
    return false, nil
  }

  _, err = REDIS_CLIENT.Del(ctx, key).Result()

  if err != nil {
    return false, err
  }

  return true, nil
}
