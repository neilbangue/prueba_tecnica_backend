package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"server/mappings"
	"server/models"
	"testing"
)

func TestServer(t *testing.T) {

	mappings.CreateUrlMappings()

	server := httptest.NewServer(mappings.Router)

	defer server.Close()

	url := fmt.Sprintf("%s/v1/messages/", server.URL)

	message := &models.Message{Text: "Testing api server"}
	requestBody, err := json.Marshal(message)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	wsResponse, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	defer wsResponse.Body.Close()

	body, err := ioutil.ReadAll(wsResponse.Body)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var response models.Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	assert.Equal(t, 18, response.Count)
}
