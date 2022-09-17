
package handlers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/0x30c4/GoPasteBin/internal/env"
	"io"
	"os"
	"log"
	"path/filepath"
	"strings"	
)

func Index(c *gin.Context) {
	// TODO Show different results for curl and browser
	c.String(http.StatusOK, "hello\n")
}

func GetData(c *gin.Context) {
	// reading the requested data from the disk and returning it
	pasteId := c.Param("ID")

	// check for LFI in the ID field
	if strings.Contains(pasteId, ".") || strings.Contains(pasteId, "/"){
		c.String(403, "Look like you attacking me")
        return
	}

	// creating path for the file
    targetPath := filepath.Join(env.DATA_DIR, pasteId)
    c.Header("Content-Type", "text/plain")
    c.File(targetPath)
}

func PostData(c *gin.Context) {

	// TODO add db functionality.
	// Store ClientIP, UserAgent, Filename if available.

	// Reading the request and creading a new file with the same uuid used as PasteID
  	file, _, err := c.Request.FormFile("f") 

  	if err != nil {
  		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
  		return
  	}

	// log.Println(c.ClientIP(), c.Request.UserAgent(), header.Filename)
	// creating UUID for every new paste
	uuidFileName := uuid.New().String()

  	out, err := os.Create(env.DATA_DIR + "/" + uuidFileName)

  	if err != nil {
  		log.Fatal(err)
  	}
	
  	defer out.Close()
  	_, err = io.Copy(out, file)
  	if err != nil {
  		log.Fatal(err)
  	}

	// returnin the url of the paste
	// remove the " from PASTE_DATA_PATH
	pastePath := strings.Replace(env.PASTE_DATA_PATH, "\"", -1)
  	filePath := filepath.Join(env.DOMAIN, , uuidFileName)
	c.String(http.StatusCreated, filePath)
}

