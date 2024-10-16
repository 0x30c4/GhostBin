package services

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/0x30c4/ghostbin/internal/repositories"
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

func TestBasicServices(t *testing.T) {
	// Setup Redis
  rdb, s := setupTestRedis(t)
  defer s.Close()

	pasteRepo := repositories.NewPasteRepository(rdb)
	pasteService := NewPasteService(pasteRepo)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a new paste
	pasteID, err := pasteService.CreateNewPaste(ctx, 60, 5, 10, "my-secret")
	if err != nil {
		fmt.Println("Failed to create paste:", err)
		return
	}

	fmt.Println("Created paste with ID:", pasteID)

	// Check if paste exists
	exists, err := pasteService.IsPasteExist(ctx, pasteID)
	if err != nil {
		fmt.Println("Failed to check paste existence:", err)
		return
	}

	fmt.Println("Paste exists:", exists)

  assert.Equal(t, 10, len(pasteID))

	fmt.Println("PasteID length is:", len(pasteID))
	// Retrieve paste
	paste, err := pasteService.GetPaste(ctx, pasteID)
	if err != nil {
		fmt.Println("Failed to get paste:", err)
		return
	}

	fmt.Println("Retrieved paste:", paste)

	// Delete paste
	deleted, err := pasteService.DeletePaste(ctx, pasteID, "my-secret")
	if err != nil {
		fmt.Println("Failed to delete paste:", err)
		return
	}

	fmt.Println("Paste deleted:", deleted)
}
