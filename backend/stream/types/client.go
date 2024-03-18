package types

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	userId     uint64
	liveConn   *websocket.Conn
	recordConn *websocket.Conn
	live       *Live
}

func (c *Client) GetUserId() uint64 {
	return c.userId
}

func (c *Client) SetLiveConn(conn *websocket.Conn) {
	c.liveConn = conn
}

func (c *Client) GetLiveConn() *websocket.Conn {
	return c.liveConn
}

func (c *Client) SetRecordConn(conn *websocket.Conn) {
	c.recordConn = conn
}

func (c *Client) GetRecordConn() *websocket.Conn {
	return c.recordConn
}

func (c *Client) SetLive(live *Live) {
	c.live = live
}

func (c *Client) GetLive() *Live {
	return c.live
}

func (c *Client) PushStream(msg []byte) error {
	return c.liveConn.WriteMessage(websocket.BinaryMessage, msg)
}

func (c *Client) ReadStream() (int, []byte, error) {
	return c.liveConn.ReadMessage()
}

func (c *Client) Write(msg []byte) error {
	return c.recordConn.WriteMessage(websocket.TextMessage, msg)
}

func (c *Client) Read() (int, []byte, error) {
	return c.recordConn.ReadMessage()
}

func (c *Client) Close() {
	if c.liveConn != nil {
		c.liveConn.Close()
	}
	if c.recordConn != nil {
		c.recordConn.Close()
	}
}
