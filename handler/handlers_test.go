package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUrl(t *testing.T) {
	requestBody, err := json.Marshal(map[string]string{
		"redirect_name": "google",
		"original_url":  "https://google.com",
	})
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "/urls/create", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUrl)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
