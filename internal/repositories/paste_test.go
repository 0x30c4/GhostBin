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


func TestCreateNewPaste(t *testing.T) {
  rdb, s := setupTestRedis(t)
  defer s.Close()

  repo := NewPasteRepository(rdb)

  ctx, cancle := context.WithTimeout(context.Background(), 2 * time.Second)
  defer cancle()

  var burnAfter uint64 = 10
  var readCount uint64 = 1
  var deepUrl uint8 = 10
  var secret string = "hello"
  var pasteID string = ""
  var pasteExist bool = false

  pasteID, err := repo.GetNewPasteID(ctx)
  assert.Nil(t, err)
  err = repo.CreatePaste(ctx, pasteID, burnAfter, readCount, deepUrl, secret)
  assert.Nil(t, err)

  pasteExist = repo.IsPasteExist(ctx, pasteID)

  assert.Equal(t, true, pasteExist)

  // read test
  for i := range readCount + 2 {
    if i <= readCount - 1 {
      exists, err := repo.GetPaste(ctx, pasteID)
      assert.Nil(t, err)
      assert.Equal(t, true, exists)
    }
    if i >= readCount - 1 {
      exists, err := repo.GetPaste(ctx, pasteID)
      assert.Nil(t, err)
      assert.Equal(t, false, exists)
    }
  }
}

func TestSecretMatch(t *testing.T) {
  rdb, s := setupTestRedis(t)
  defer s.Close()

  repo := NewPasteRepository(rdb)

  ctx, cancle := context.WithTimeout(context.Background(), 2 * time.Second)
  defer cancle()

  var burnAfter uint64 = 10
  var readCount uint64 = 1
  var deepUrl uint8 = 10
  var secret string = "hello"
  var pasteID string = ""
  var pasteExist bool = false

  pasteID, err := repo.GetNewPasteID(ctx)
  assert.Nil(t, err)
  err = repo.CreatePaste(ctx, pasteID, burnAfter, readCount, deepUrl, secret)
  assert.Nil(t, err)

  pasteExist = repo.IsPasteExist(ctx, pasteID)

  assert.Equal(t, true, pasteExist)

  ok, err := repo.SecretMatch(ctx, pasteID, secret)

  assert.Nil(t, err)
  assert.Equal(t, true, ok)

  ok, err = repo.SecretMatch(ctx, pasteID, secret + "1")

  assert.Nil(t, err)
  assert.Equal(t, false, ok)
}

func TestDeletePaste(t *testing.T) {
  rdb, s := setupTestRedis(t)
  defer s.Close()

  repo := NewPasteRepository(rdb)

  ctx, cancle := context.WithTimeout(context.Background(), 2 * time.Second)
  defer cancle()

  var burnAfter uint64 = 10
  var readCount uint64 = 1
  var deepUrl uint8 = 10
  var secret string = "hello"
  var pasteID string = ""
  var pasteExist bool = false

  pasteID, err := repo.GetNewPasteID(ctx)
  assert.Nil(t, err)
  err = repo.CreatePaste(ctx, pasteID, burnAfter, readCount, deepUrl, secret)
  assert.Nil(t, err)

  pasteExist = repo.IsPasteExist(ctx, pasteID)

  assert.Equal(t, true, pasteExist)

  deleteDone, err := repo.DeletePaste(ctx, pasteID, secret)

  assert.Nil(t, err)
  assert.Equal(t, true, deleteDone)

}
