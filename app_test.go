package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"gourlshort"
	"gourlshort/model"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}

	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	_, err := a.DB.Exec("TRUNCATE TABLE urls")
	if err != nil {
		log.Fatalln(err)
	}
}

const (
	tableCreationQuery = `CREATE TABLE IF NOT EXISTS urls
(
  id INT AUTO_INCREMENT PRIMARY KEY,
  redirect_name varchar(55) NOT NULL UNIQUE,
  original_url varchar(55) NOT NULL UNIQUE
)`
)

func TestCreateUrlInvalidPayload(t *testing.T) {
	clearTable()

	requestBody, err := json.Marshal(map[string]int{
		"name": 1,
	})
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "/urls/create", bytes.NewBuffer(requestBody))
	fmt.Println()
	response := executeRequest(req)

	fmt.Println(response.Code)
	checkResponseCode(t, http.StatusBadRequest, response.Code)

	errorMessage := "Invalid request payload"
	responseError := getResponseError(response)
	if responseError != errorMessage {
		t.Errorf("Expected, this error: %s, got: %s", errorMessage, responseError)
	}

}

func TestCreateUrl(t *testing.T) {
	clearTable()

	requestBody, err := json.Marshal(map[string]string{
		"redirect_name": "google",
		"original_url":  "https://google.com",
	})
	if err != nil {
		log.Fatalln(err)

	}

	req, err := http.NewRequest("POST", "/urls/create", bytes.NewBuffer(requestBody))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var url model.URL
	err = json.Unmarshal(response.Body.Bytes(), &url)
	if err != nil {
		log.Fatalln(err)
	}

	if url.ID != 0 {
		t.Errorf("Expected id of the url to be '0', got '%v'", url.ID)
	}
	if url.RedirectName != "google" {
		t.Errorf("Expected redirect name to be 'google', got '%v'", url.RedirectName)
	}
	if url.OriginalUrl != "https://google.com" {
		t.Errorf("Expected original url to be 'https://google.com', got '%v'", url.OriginalUrl)
	}
}

func TestCreateDuplicate(t *testing.T) {
	clearTable()
	addUrlToDb("yahoo", "https://yahoo.com")

	requestBody, err := json.Marshal(map[string]string{
		"redirect_name": "yahoo",
		"original_url":  "https://yahoo.com",
	})
	if err != nil {
		log.Fatalln(err)

	}


	req, err := http.NewRequest("POST", "/urls/create", bytes.NewBuffer(requestBody))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotAcceptable, response.Code)

	errorMessage := "Duplicate url found"
	responseError := getResponseError(response)

	if responseError != errorMessage {
		t.Errorf("Expected, this error: %s, got: %s", errorMessage, responseError)
	}
}

func TestGetNonExistentUrl(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/urls/google", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestGetUrl(t *testing.T) {
	clearTable()
	addUrlToDb("google", "https://google.com")

	req, err := http.NewRequest("GET", "/urls/google", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)

	checkResponseCode(t, http.StatusFound, response.Code)
}

func addUrlToDb(redirectName string, originalUrl string) {
	_, err := a.DB.Exec("INSERT INTO urls(redirect_name, original_url) VALUES (?, ?)", redirectName, originalUrl)
	if err != nil {
		log.Fatalln(err)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func getResponseError(resp *httptest.ResponseRecorder) string {
	var data map[string]string
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}

	return data["error"]
}
