package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"k8s.io/apimachinery/pkg/util/uuid"
)

type WsConnection struct {
	*websocket.Conn
	id    string
	read  chan *WsMessage
	write chan interface{}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWsConnection(c *gin.Context) (*WsConnection, error) {
	connection, err := upgrader.Upgrade(
		c.Writer,
		c.Request,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &WsConnection{
		id:    string(uuid.NewUUID()),
		Conn:  connection,
		read:  make(chan *WsMessage),
		write: make(chan interface{}),
	}, nil
}

func (wc *WsConnection) GetID() string {
	return wc.id
}

func (wc *WsConnection) receive(callback func(key string)) {
	defer close(wc.read)
	defer func() {
		callback(wc.id)
	}()

	for {
		_, data, err := wc.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := &WsMessage{}
		err = json.Unmarshal(data, message)

		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(message)
			wc.read <- message
		}
	}
}

func (wc *WsConnection) loop(callback func(wm *WsMessage) interface{}) {
	defer close(wc.write)

	for msg := range wc.read {
		if callback != nil {
			wc.write <- callback(msg)
		} else {
			wc.write <- msg
		}
	}
}

func (wc *WsConnection) send() {

	for msg := range wc.write {
		err := wc.WriteJSON(msg)
		if err != nil {
			log.Println(err)
		}
	}
}
