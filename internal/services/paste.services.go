package services

import (
	"context"
	"errors"
	"github.com/0x30c4/ghostbin/internal/repositories"
)

type PasteService struct {
	pasteRepo *repositories.PasteRepository
}

// NewPasteService creates a new instance of the PasteService
func NewPasteService(pasteRepo *repositories.PasteRepository) *PasteService {
	return &PasteService{pasteRepo: pasteRepo}
}

// CreateNewPaste creates a new paste and returns the paste ID
func (s *PasteService) CreateNewPaste(ctx context.Context, burnAfter, readCount uint64, deepUrl uint8, secret string) (string, error) {

	pasteID, err := s.pasteRepo.GetNewPasteID(ctx)
	if err != nil {
		return pasteID, err
	}

  if readCount == 0 {
    readCount = 4096
  }

  if burnAfter == 0 {
    burnAfter = 5961600
  }

  if deepUrl >= 8 {
    var longPasteId string
    var checkStatus bool
    for i := 0; i <= 8; i++ {
      longPasteId = randomPasteIdPrefix(deepUrl)
      checkStatus, err = s.IsPasteExist(ctx, longPasteId)
      if !checkStatus {
        pasteID = longPasteId
        break
      }
      if i == 8 && checkStatus {
        pasteID = ""
        return pasteID, err
      }
    }
  }

	err = s.pasteRepo.CreatePaste(ctx, pasteID, burnAfter, readCount, deepUrl, secret)
	if err != nil {
		return "", err
	}

	return pasteID, nil
}

// IsPasteExist checks if a paste with the given ID exists
func (s *PasteService) IsPasteExist(ctx context.Context, pasteID string) (bool, error) {
	if pasteID == "" {
		return false, errors.New("pasteID cannot be empty")
	}

	return s.pasteRepo.IsPasteExist(ctx, pasteID), nil
}

// GetPaste retrieves a paste by ID, decrements its read count, and deletes it if the read count is exhausted
func (s *PasteService) GetPaste(ctx context.Context, pasteID string) (bool, error) {

  var exists bool = false
	if pasteID == "" {
		return exists, errors.New("pasteID cannot be empty")
	}

	exists, err := s.pasteRepo.GetPaste(ctx, pasteID)
	if err != nil {
		return exists, err
	}

	if !exists {
		return exists, errors.New("paste not found or expired")
	}

	return exists, err
}

func (s *PasteService) SecretMatch(ctx context.Context, pasteID string, secret string) (bool, error) {
	return s.pasteRepo.SecretMatch(ctx, pasteID, secret)
}

// DeletePaste deletes a paste if the secret matches
func (s *PasteService) DeletePaste(ctx context.Context, pasteID string, secret string) (bool, error) {
	if pasteID == "" || secret == "" {
		return false, errors.New("pasteID and secret cannot be empty")
	}
	return s.pasteRepo.DeletePaste(ctx, pasteID, secret)
}
