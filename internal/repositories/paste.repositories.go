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

func (r *PasteRepository) CreatePaste(pasteModel models.Paste) error {
  return nil
}

func (r *PasteRepository) IsPasteExist(pasteId string) bool {
  return false
}

func (r *PasteRepository) GetPaste(pasteId string) (bool, error) {
  return false, nil
}

func (r *PasteRepository) DeletePaste(pasteId string, secret string) (bool, error) {
  return false, nil
}
