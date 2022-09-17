package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/0x30c4/GoPasteBin/api/v1/router"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"bytes"
	"path"
	"path/filepath"
	"log"
	"strings"
)

func TestIndexHandler(t *testing.T) {

	// TODO show different results for curl and browser request
	router := router.Initialize()

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello\n", w.Body.String())
}

func TestPostDataHandler(t *testing.T) {
	router := router.Initialize()

	w := httptest.NewRecorder()


	// Reading the test file and posting it to the endpoint
	fileDir, _ := os.Getwd()
  	fileName := "handlers_test.go" // the test file name
  	filePath := path.Join(fileDir, fileName)

  	file, _ := os.Open(filePath)
  	defer file.Close()

  	body := &bytes.Buffer{}
  	writer := multipart.NewWriter(body)
  	part, _ := writer.CreateFormFile("f", filepath.Base(file.Name()))
  	io.Copy(part, file)
  	writer.Close()

	// creating request 
  	r := httptest.NewRequest(http.MethodPost, "/", body)

  	r.Header.Add("Content-Type", writer.FormDataContentType())

	router.ServeHTTP(w, r)

	// check status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// check if the returned url is valid
	bodyBytes, err := io.ReadAll(w.Body)
    if err != nil {
        log.Fatal(err)
    }
    bodyString := string(bodyBytes)

	log.Println(bodyString)
	pasteId := strings.Split(bodyString, "/")[2]

	// create a new request to fetch the newly created file
	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/" + pasteId, nil)

	router.ServeHTTP(w, r)
	// reading the original file to compare it to the fetched data
	originalFile, err := ioutil.ReadFile(fileName) // just pass the file name
    if err != nil {
        log.Fatal(err)
    }
	bodyBytes, err = io.ReadAll(w.Body)
    if err != nil {
        log.Fatal(err)
    }
	
	bodyString = string(bodyBytes)

	originalFileStr := string(originalFile)

	// Checking if the original file and the fetched data is same or not.
	assert.Equal(t, originalFileStr, bodyString)

}
