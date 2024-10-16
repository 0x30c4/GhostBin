package services

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/0x30c4/ghostbin/internal/repositories"
	"github.com/0x30c4/ghostbin/internal/utils"
)

type FileService struct {
	pasteRepo   *repositories.PasteRepository
	uploadDir   string
	maxFileSize uint64
}

func NewFileService(pasteRepo *repositories.PasteRepository, uploadDir, maxFileSize string) *FileService {
  var maxFileSizeUINT64 uint64

  maxFileSizeUINT64, err := utils.ParseSizeToBytes(maxFileSize)

  if err != nil {
    maxFileSizeUINT64 = 2 * 1024 * 1024 // default of 2MB
  }


	return &FileService{
    pasteRepo: pasteRepo,
    uploadDir: uploadDir,
    maxFileSize: maxFileSizeUINT64,
  }
}

// UploadFile validates and stores a file, then saves metadata to Redis
func (s *FileService) SaveFile(pasteID string, fileSize uint64, file io.Reader) (string, error) {

  var filePath string

  if !s.isFileSizeOK(fileSize) {
    return filePath, fmt.Errorf("File size exceeded the limit of %d < %d ", s.maxFileSize, fileSize)
  }

  // Create a new file on the server
  filePath = path.Join(s.uploadDir, pasteID)

  dst, err := os.Create(filePath)
  if err != nil {
    return filePath, err
  }

  defer dst.Close()

  // Copy the uploaded file data to the newly created file
  _, err = io.Copy(dst, file)
  if err != nil {
    return filePath, err
  }

	return filePath, nil
}

func (s *FileService) isFileSizeOK(fileSize uint64) bool {
  return fileSize <= s.maxFileSize
}
