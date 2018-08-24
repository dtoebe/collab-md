package hub

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMsgSize = 512
)

var (
	newLine  = []byte{'\n'}
	space    = []byte{' '}
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type client struct {
	conn *websocket.Conn
	send chan *message
	key  string
}

type message struct {
	broadcast []byte
	senderKey string
}

// readPump is the message receiver
// func (c *client) readPump() {
// 	defer c.conn.Close()

// 	c.conn.SetReadLimit(maxMsgSize)
// 	c.conn.SetReadDeadline(time.Now().Add(pongWait))
// 	c.conn.SetPongHandler(func(string) error {
// 		c.conn.SetReadDeadline(time.Now().Add(pongWait))
// 		return nil
// 	})

// 	_, msg, err := c.conn.ReadMessage()
// 	if err != nil {
// 		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
// 			log.Println(err)
// 		}
// 		break
// 	}

// }
