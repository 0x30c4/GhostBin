package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"github.com/0x30c4/ghostbin/internal/env"
	"github.com/0x30c4/ghostbin/internal/redis"
)

func Index(w http.ResponseWriter, r *http.Request){
  userAgent := r.Header.Get("User-Agent")
  w.Header().Set("Content-Type", "text/html; charset=utf-8")

  if CheckIfBrowser(userAgent) {
    // Open the file
    file, err := os.Open("./public/index.html")
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to open file: %s", err), http.StatusInternalServerError)
        return
    }
    defer file.Close()

  // Copy the file content to the response writer
  _, err = io.Copy(w, file)
  if err != nil {
      http.Error(w, fmt.Sprintf("Failed to copy file content: %s", err), http.StatusInternalServerError)
      return
  }

  }else {
    fmt.Fprintf(w, "TL;DR: $ curl -F \"f=@filename.ext\" %s \n", env.DOMAIN)
    fmt.Fprintf(w, "TL;DR: $ cat file | curl -F \"f=@-\" %s \n", env.DOMAIN)
  }
}


func GetPaste(w http.ResponseWriter, r *http.Request) {
  pasteId := r.PathValue("id")

  if CheckPaste(pasteId) {
    ok, err := redis.GetPasteRDB(pasteId)

    if err != nil {
      log.Printf("Server Error [%s]\n", err)
      http.Error(w, "Internal Server Error", http.StatusInternalServerError)
      return
    }
    if !ok {
      http.Error(w, "Not Found", http.StatusNotFound)
      return
    }

    ReadFile(pasteId, w)
  }else {
    errorMessage := fmt.Sprintf("Paste ID [%s] doesn't exists", pasteId)
    http.Error(w, errorMessage, http.StatusInternalServerError)
  }
}

func NewPaste(w http.ResponseWriter, r *http.Request) {

  // Parse the multipart form
  const maxUploadSize = 2 * 1024 * 1024

  if r.ContentLength > maxUploadSize {
    log.Printf("Server Error [File size bigger than 2MB]\n")
    http.Error(w, "File size bigger than 2MB", http.StatusBadRequest)
    return
  }

  // Retrieve the file from the form data
  file, _, err := r.FormFile("f")

  if err != nil {
    log.Printf("Server Error [%s]\n", err)
    http.Error(w, "No file provided", http.StatusBadRequest)
    return
  }
  defer file.Close()

  var pasteModel redis.PasteModel


  pasteModel.PasteId, err = redis.GetNewPasteID()

  if err != nil {
    log.Printf("Server Error [%s]\n", err)
    http.Error(w, "Failed to create file on the server", http.StatusInternalServerError)
    return
  }

  expire := r.FormValue("expire")
  readCount := r.FormValue("read")

  deepUrl := r.FormValue("deepurl")

  pasteModel.Secret = r.FormValue("secret")

  // convert expire after readCount value

  StrToInt(deepUrl, &pasteModel.DeepUrl, 8)

  StrToInt(expire, &pasteModel.BurnAfter, 64)

  StrToInt(readCount, &pasteModel.ReadCount, 64)

  if pasteModel.ReadCount == 0 {
    pasteModel.ReadCount = 4096
  }

  if pasteModel.BurnAfter == 0 {
    pasteModel.BurnAfter = 5961600
  }

  if pasteModel.DeepUrl >= 8 {
    var longPasteId string
    var checkStatus bool
    for i := 0; i <= 8; i++ {
      longPasteId = RandomPasteIdPrefix(pasteModel.DeepUrl)
      checkStatus = redis.IsPasteExist(longPasteId)
      if !checkStatus {
        pasteModel.PasteId = longPasteId
        break
      }
      if i == 8 && checkStatus {
        log.Printf("Can't generate an uniq unique pasteId [%s]\n", err)
        http.Error(w, "Failed try again", http.StatusInternalServerError)
        return
      }
    }
  }

  err = redis.PutPasteRDB(pasteModel)

  if err != nil {
    log.Printf("Server Error [%s]\n", err)
    http.Error(w, "Failed to create file on server", http.StatusInternalServerError)
    return
  }

  // Create a new file on the server
  filePath := path.Join(env.PASTE_DIR, pasteModel.PasteId)

  dst, err := os.Create(filePath)
  if err != nil {
    log.Printf("Server Error [%s]\n", err)
    http.Error(w, "Failed to create file on server", http.StatusInternalServerError)
    return
  }

  defer dst.Close()

  // Copy the uploaded file data to the newly created file
  _, err = io.Copy(dst, file)
  if err != nil {
    log.Printf("Server Error [%s]\n", err)
    http.Error(w, "Failed to copy file to server", http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "%s/%s", env.DOMAIN, pasteModel.PasteId)

}

func EditPaste(w http.ResponseWriter, r *http.Request) {
  // TODO: implement EditPaste
  // pasteId := r.PathValue("id")
  // secret := r.FormValue("secret")

}

func DeletePaste(w http.ResponseWriter, r *http.Request) {

  pasteId := r.PathValue("id")

  secret := r.FormValue("secret")

  ok, err := redis.DeletePasteRDB(pasteId, secret)

  if err != nil {
    log.Println(err)
    http.Error(w, "Failed to delete the file from the server.", http.StatusInternalServerError)
    return
  }

  if !ok {
    http.Error(w, "Secret doesn't match", http.StatusUnauthorized)
    return
  }

  fmt.Fprintf(w, "%s was deleted.", pasteId)
}
