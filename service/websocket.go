package service

import (
	"errors"
	"kallme/config"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conns   = make(map[string]*websocket.Conn)
	message = make(map[string]chan interface{})
	mux     sync.Mutex
)

func SendMessage(token string, msg interface{}) (err error) {
	_, exist := getMessageChan(token)
	err = nil
	if !exist {
		config.Log.Error("Message channel not exist")
		err = errors.New("message channel not exist")
		return
	}

	setMessage(token, msg)

	return
}

func HandleWebSocket(c *gin.Context, token string) {
	var (
		conn  *websocket.Conn
		err   error
		exist bool
		m     chan interface{}
	)

	conn, err = upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		config.Log.Error("Failed to upgrade connection:", err)
		return
	}

	addConnection(token, conn)

	m, exist = getMessageChan(token)
	if !exist {
		m = make(chan interface{})
		addMessageChan(token, m)
	}
	config.Log.Info("connect success")

	for {
		select {
		case msg, _ := <-m:
			err = conn.WriteJSON(msg)
			config.Log.Info("Message: ", msg)
			if err != nil {
				config.Log.Error("Failed to write message:", err)
				conn.Close()
				return
			}
		}
		// messageType, message, err := conn.ReadMessage()
		// if err != nil {
		// 	config.Log.Error("Failed to read message:", err)
		// 	break
		// }
		// config.Log.Info("Received message: ", message)

		// err = conn.WriteMessage(messageType, message)
		// if err != nil {
		// 	config.Log.Error("Failed to write message:", err)
		// 	break
		// }
	}

}

func addConnection(token string, conn *websocket.Conn) {
	mux.Lock()
	defer mux.Unlock()
	conns[token] = conn
}

func getConnection(token string) (conn *websocket.Conn, ok bool) {
	mux.Lock()
	defer mux.Unlock()
	conn, ok = conns[token]
	return
}

func delectConnection(token string) {
	mux.Lock()
	defer mux.Unlock()
	delete(conns, token)
}

func addMessageChan(token string, msg chan interface{}) {
	mux.Lock()
	defer mux.Unlock()
	message[token] = msg
}

func getMessageChan(token string) (msg chan interface{}, ok bool) {
	mux.Lock()
	defer mux.Unlock()
	msg, ok = message[token]
	return
}

func delectMessageChan(token string) {
	mux.Lock()
	defer mux.Unlock()
	if msg, ok := message[token]; ok {
		close(msg)
		delete(message, token)
	}
}

func setMessage(id string, content interface{}) {
	mux.Lock()
	if m, exist := message[id]; exist {
		go func() {
			m <- content
		}()
	}
	mux.Unlock()
}
