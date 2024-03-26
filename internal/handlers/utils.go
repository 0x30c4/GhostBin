package handlers

import (
  "io"
	"os"
  "fmt"
  "log"
	"path"
  "time"
	"strings"
  "reflect"
	"strconv"
  "net/http"
  "math/rand"
	"github.com/0x30c4/ghostbin/internal/env"
)

func CheckIfBrowser(userAgent string) bool {

  if strings.Contains(userAgent, "Mozilla") || strings.Contains(userAgent, "Chrome") || strings.Contains(userAgent, "Safari") {
    // Client is likely a web browser
    return true
  } else {
    // Client may not be a web browser
    return false
  }
}

func CheckPaste(pasteId string) bool {
  filename := path.Join(env.PASTE_DIR, pasteId)
  // Check if the file exists
  if _, err := os.Stat(filename); os.IsNotExist(err) {
    return false
  } else {
    return true
  }
}


func ReadFile(pasteId string, w http.ResponseWriter) {

  filename := path.Join(env.PASTE_DIR, pasteId)
  // Open the file
  file, err := os.Open(filename)
  if err != nil {
      http.Error(w, fmt.Sprintf("Failed to open file: %s", err), http.StatusInternalServerError)
      return
  }
  defer file.Close()

  // Set the appropriate content type
  w.Header().Set("Content-Type", "text/plain")

  // Copy the file content to the response writer
  _, err = io.Copy(w, file)
  if err != nil {
      http.Error(w, fmt.Sprintf("Failed to copy file content: %s", err), http.StatusInternalServerError)
      return
  }
}

// Generate a random string of specified length
func RandomPasteIdPrefix(length uint8) string {
	// Define the character set
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a byte slice of the specified length
	randomString := make([]byte, length)

	// Fill the byte slice with random characters from the character set
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}

	// Convert the byte slice to a string and return it
	return string(randomString)
}

func StrToInt[T interface{}](str string, valueNum *T, size int) {
    num, err := strconv.ParseUint(str, 10, size)
    if err != nil {
        log.Println("Error:", err)
    } else {
        value := reflect.ValueOf(valueNum).Elem()
        switch value.Kind() {
        case reflect.Uint8:
            value.SetUint(uint64(num))
        case reflect.Uint64:
            value.SetUint(num)
        default:
            log.Println("Unsupported type")
        }
    }
}


  // if readCount != "" {
  //   num, err := strconv.ParseUint(readCount, 10, 64)
  //   if err != nil {
  //     log.Println("Error:", err)
  //   } else {
  //     pasteModel.ReadCount = num
  //   }
  // }
