package controllers

import (
	"connectity/websocket/models"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Handler(writer http.ResponseWriter, request *http.Request) {
	ws, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		println(err.Error())
		return
	}

	for {
		var message models.Message
		err := ws.ReadJSON(&message)
		if err != nil {
			println(err.Error())
			return
		}

		fmt.Printf("client> %s\n", message.Text)

		response := &models.Response{Count: len(message.Text)}

		fmt.Printf("server> %d\n", response.Count)

		if err := ws.WriteJSON(response); err != nil {
			println(err.Error())
			return
		}
	}
}
