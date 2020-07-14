package main

import (
	"connectity/websocket/controllers"
	"connectity/websocket/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controllers.Handler))
	defer server.Close()

	url := "ws" + strings.TrimPrefix(server.URL, "http")

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	message := &models.Message{Text: "Testing websocket"}

	if err := ws.WriteJSON(message); err != nil {
		t.Fatalf("%v", err)
	}

	var response models.Response

	if err := ws.ReadJSON(&response); err != nil {
		t.Fatalf("%v", err)
	}

	assert.Equal(t, 17, response.Count)
}
