package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/0x30c4/ghostbin/pkg/utils"
)

const (
  plainTextContentType = "text/plain"
  htmlContentType      = "text/html; charset=utf-8"
)

var (
  indexFile = os.Getenv("INDEX_FILE")
  indexFileCLI = os.Getenv("INDEX_FILE_CLI")
  domain = os.Getenv("DOMAIN")

)


func (h *HttpHandlers) Index(w http.ResponseWriter, r *http.Request){
  userAgent := r.Header.Get("User-Agent")
  var sendIndex string = indexFileCLI
  if checkIfBrowser(userAgent) {
    sendIndex = indexFile
  }

  responseWritten, err := readFile(sendIndex, w)

  if err != nil {
    h.logger.Error(err.Error())
    http.Error(w, "Server Error", http.StatusInternalServerError)
    return
  }

  // Set the appropriate content type
  setContentType(htmlContentType, w)

  h.logger.Info("Response Written",
    slog.Int64("bytes", responseWritten),
  )
}

func (h *HttpHandlers) GetPaste(w http.ResponseWriter, r *http.Request) {
  pasteID := r.PathValue("id")

  setContentType(plainTextContentType, w)

  if !checkPasteInFS(pasteID) {
    h.logger.Info("Paste Not not found on the system",
      slog.String("pasteID", pasteID),
    )
    http.Error(w, "Not Found", http.StatusNotFound)
    return
  }

  ok, err := h.pasteSrv.GetPaste(context.Background(), pasteID)

  if err != nil {
    h.logger.Info("PasteService error",
      slog.String("pasteID", pasteID),
      slog.String("err", err.Error()),
    )
    http.Error(w, "Not Found", http.StatusNotFound)
    return
  }

  if !ok {
    h.logger.Info("Paste Not not found",
      slog.String("pasteID", pasteID),
    )
    http.Error(w, "Not Found", http.StatusNotFound)
    return
  }

  readFile(pasteID2Path(pasteID), w)
}

func (h *HttpHandlers) NewPaste(w http.ResponseWriter, r *http.Request) {

  // Retrieve the file from the form data
  file, _, err := r.FormFile("f")

  if err != nil {
    h.logger.Info("Server Error [%s]\n", err)
    http.Error(w, "No file provided", http.StatusBadRequest)
    return
  }
  defer file.Close()

  expireStr := r.FormValue("expire")
  readCountStr := r.FormValue("read")
  deepUrlStr := r.FormValue("deepurl")
  secret := r.FormValue("secret")

  var deepUrl uint8
  var expire uint64
  var readCount uint64

  utils.StrToInt(deepUrlStr, &deepUrl, 8)

  utils.StrToInt(expireStr, &expire, 64)

  utils.StrToInt(readCountStr, &readCount, 64)

  pasteID, err := h.pasteSrv.CreateNewPaste(context.Background(), expire, readCount, deepUrl, secret)

  if err != nil {
    h.logger.Info("Server Error [%s]\n", err)
    http.Error(w, "Failed to create file on the server", http.StatusInternalServerError)
    return
  }

  h.fileSrv.SaveFile(pasteID, uint64(r.ContentLength), file)

  fmt.Fprintf(w, "%s/%s\n", domain, pasteID)

}

// func EditPaste(w http.ResponseWriter, r *http.Request) {
//   // TODO: implement EditPaste
//   // pasteId := r.PathValue("id")
//   // secret := r.FormValue("secret")
// }

func (h *HttpHandlers) DeletePaste(w http.ResponseWriter, r *http.Request) {

  pasteID := r.PathValue("id")

  secret := r.FormValue("secret")

  ok, err := h.pasteSrv.DeletePaste(context.Background(), pasteID, secret)

  if err != nil {
    h.logger.Info("PasteService error",
      slog.String("err", err.Error()),
    )
    http.Error(w, "Failed to delete the file from the server.", http.StatusInternalServerError)
    return
  }

  if !ok {
    h.logger.Info("Secret doesn't match",
      slog.String("pasteID", pasteID),
    )
    http.Error(w, "Secret doesn't match", http.StatusUnauthorized)
    return
  }

  fmt.Fprintf(w, "%s was deleted.", pasteID)
}
