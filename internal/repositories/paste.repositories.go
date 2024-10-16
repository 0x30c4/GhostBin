package repositories

import (
	"context"
	"github.com/0x30c4/ghostbin/internal/models"
	"github.com/0x30c4/ghostbin/internal/utils"
	"github.com/redis/go-redis/v9"
)

type PasteRepository struct {
  redisClient *redis.Client
}

// Create a new instance of the PasteRepository
func NewPasteRepository(redisClient *redis.Client) *PasteRepository {
  return &PasteRepository{redisClient: redisClient}
}

func (r *PasteRepository) GetNewPasteID(ctx context.Context) (string, error) {

  tx := r.redisClient.TxPipeline()

  var newPasteID string = ""

  tx.Incr(ctx, "totalPaste")
  totalPaste := tx.Get(ctx, "totalPaste")

  _, err := tx.Exec(ctx)
  if err != nil {
    return newPasteID, err
  }

  totalPasteValue, err := totalPaste.Int64()

  if err != nil {
    return newPasteID, err
  }

  newPasteID = utils.HexStr(totalPasteValue)

  return newPasteID, nil
}

func (r *PasteRepository) CreatePaste(ctx context.Context, pasteID string, burnAfter, readCount uint64, deepUrl uint8, secret string) error {

  var pasteModel = models.Paste{
    PasteId: pasteID,
    BurnAfter: burnAfter,
    ReadCount: readCount,
    DeepUrl: deepUrl,
    Secret: secret,
  }

  if pasteModel.ReadCount == 0 {
    return nil
  }

  key := "pasteId:" + pasteModel.PasteId

  err := r.redisClient.HSet(ctx, key, pasteModel).Err()

  if err != nil {
    return err
  }

  err = r.redisClient.Do(ctx, "EXPIRE", key, pasteModel.BurnAfter).Err()

  return err
}

func (r *PasteRepository) IsPasteExist(ctx context.Context, pasteId string) bool {
  key := "pasteId:" + pasteId
  exists, err := r.redisClient.Exists(ctx, key).Result()
  if err != nil {
    return false
  }

  if exists > 0 {
    return true
  } else {
    return false
  }
}

func (r *PasteRepository) GetPaste(ctx context.Context, pasteId string) (bool, error) {
  key := "pasteId:" + pasteId
  res := r.redisClient.HGetAll(ctx, key)

  err := res.Err()
  if err != nil {
		return false, err
	}

  var pasteModel models.Paste

  err = res.Scan(&pasteModel)

  if (models.Paste{}) == pasteModel {
    return false, nil
  }

	if err != nil {
    return false, err
	}

  if pasteModel.ReadCount <= 1 {
    _ = r.redisClient.Del(ctx, key)
    return true, nil
  }
  err = r.redisClient.HIncrBy(ctx, key, "ReadCount", -1).Err()

  if err != nil {
    return false, nil
  }

  return true, nil
}


func (r *PasteRepository) SecretMatch(ctx context.Context, pasteId string, secret string) (bool, error) {
  key := "pasteId:" + pasteId

  secretDB, err := r.redisClient.HGet(ctx, key, "Secret").Result()

  if err != nil {
    return false, err
  }

  if len(secretDB) == 0 {
    return false, nil
  }

  if secretDB != secret {
    return false, nil
  }
  return true, nil
}

func (r *PasteRepository) DeletePaste(ctx context.Context, pasteId string, secret string) (bool, error) {
  key := "pasteId:" + pasteId

  if ok, err := r.SecretMatch(ctx, pasteId, secret); err != nil || !ok {
    return false, err
  }

  _, err := r.redisClient.Del(ctx, key).Result()

  if err != nil {
    return false, err
  }

  return true, nil
}
