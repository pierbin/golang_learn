package ws

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// TextApi webSocket returns text format
func TextApi(c *gin.Context) {
	// Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}

	defer func(ws *websocket.Conn) {
		err1 := ws.Close()
		if err1 != nil {
			return
		}
	}(ws)

	// Read data in ws
	mt, message, err2 := ws.ReadMessage()
	if err2 != nil {
		log.Println("error read message")
		log.Fatal(err2)
	}

	// Write ws data, pong 10 times
	var count = 0
	for {
		count++
		if count > 10 {
			break
		}

		message = []byte(string(message) + " " + strconv.Itoa(count))
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("error write message: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

// JsonApi webSocket returns json format
func JsonApi(c *gin.Context) {
	// Upgrade get request to webSocket protocol
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error get connection")
		log.Fatal(err)
	}

	defer func(ws *websocket.Conn) {
		err1 := ws.Close()
		if err1 != nil {
			return
		}
	}(ws)

	var data struct {
		A string `json:"a"`
		B int    `json:"b"`
	}

	// Read data in ws
	err = ws.ReadJSON(&data)
	if err != nil {
		log.Println("error read json")
		log.Fatal(err)
	}

	// Write ws data, pong 10 times
	var count = 0
	for {
		count++
		if count > 10 {
			break
		}

		err = ws.WriteJSON(struct {
			A string `json:"a"`
			B int    `json:"b"`
			C int    `json:"c"`
		}{
			A: data.A,
			B: data.B,
			C: count,
		})
		if err != nil {
			log.Println("error write json: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}
