package services

import (
	"strings"
	"testing"

	"github.com/0x30c4/ghostbin/internal/repositories"
)

func TestFileUpload(t *testing.T) {

  rdb, s := setupTestRedis(t)
  defer s.Close()

  uploadDir := "../../data/paste_data_dev"

  pasteRepo := repositories.NewPasteRepository(rdb)

  fileSrv := NewFileService(pasteRepo, uploadDir, "1212")

  str := "hello\nsdssdasdasd"

  file := strings.NewReader(str)

  fileSrv.SaveFile("hello", 12, file)
}
