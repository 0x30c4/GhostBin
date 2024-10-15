package repositories

import (
	"context"
	"testing"
	"time"

	"github.com/0x30c4/ghostbin/internal/utils"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func setupTestRedis(t *testing.T) (*redis.Client, *miniredis.Miniredis) {
  // Start a mini Redis server
  s, err := miniredis.Run()
  if err != nil {
    t.Fatalf("an error occurred: %v", err)
  }

  // Create a Redis client for testing
  rdb := redis.NewClient(&redis.Options{
    Addr: s.Addr(),
  })

  return rdb, s
}

func TestGetNewPasteID(t *testing.T) {
  rdb, s := setupTestRedis(t)
  defer s.Close()

  repo := NewPasteRepository(rdb)

  ctx, cancle := context.WithTimeout(context.Background(), 2 * time.Second)
  defer cancle()

  for i := range 1000 {
    i++
    pasteID, err := repo.GetNewPasteID(ctx)
    assert.Nil(t, err)
    assert.Equal(t, utils.HexStr(int64(i)), pasteID)
  }
}
