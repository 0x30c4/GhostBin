package handlers

import (
  "io"
	"os"
  "fmt"
	"path"
	"strings"
  "net/http"
)

var (
  uploadDir = os.Getenv("PASTE_UPLOAD_DIR")
)

func checkIfBrowser(userAgent string) bool {
  return strings.Contains(userAgent, "Mozilla") || strings.Contains(userAgent, "Chrome") || strings.Contains(userAgent, "Safari")
}

func checkPasteInFS(pasteId string) bool {
  filename := pasteID2Path(pasteId)
  // Check if the file exists
  _, err := os.Stat(filename)
  return !os.IsNotExist(err) // if not exist error true then return false
}

func pasteID2Path(pasteID string) string {
  return path.Join(uploadDir, pasteID)
}

func setContentType(contentType string, w http.ResponseWriter) {
  w.Header().Set("Content-Type", contentType)
}

func readFile(filePath string, w http.ResponseWriter) (int64, error) {

  var responseWritten int64

  // Open the file
  file, err := os.Open(filePath)
  if err != nil {
    return responseWritten, fmt.Errorf("Failed to open file: %s", err)
  }
  defer file.Close()

  // Copy the file content to the response writer
  responseWritten, err = io.Copy(w, file)
  if err != nil {
    return responseWritten, fmt.Errorf("Failed to copy file content: %s", err)
  }

  return responseWritten, nil
}
