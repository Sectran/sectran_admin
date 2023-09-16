package conn

import (
	"net"
	"time"

	"github.com/gorilla/websocket"
	"github.com/gravitational/trace"
)

type WebsocketIO struct {
	Conn      *websocket.Conn
	remaining []byte
}

func (ws *WebsocketIO) Write(p []byte) (int, error) {
	err := ws.Conn.WriteMessage(websocket.BinaryMessage, p)
	if err != nil {
		return 0, trace.Wrap(err)
	}

	return len(p), nil
}

func (ws *WebsocketIO) Read(p []byte) (int, error) {
	if len(ws.remaining) == 0 {
		ty, data, err := ws.Conn.ReadMessage()
		if err != nil {
			return 0, trace.Wrap(err)
		}
		if ty != websocket.BinaryMessage {
			return 0, trace.BadParameter("expected websocket message of type BinaryMessage, got %T", ty)
		}
		ws.remaining = data
	}

	copied := copy(p, ws.remaining)
	ws.remaining = ws.remaining[copied:]
	return copied, nil
}

func (ws *WebsocketIO) Close() error {
	return trace.Wrap(ws.Conn.Close())
}

func (ws *WebsocketIO) LocalAddr() net.Addr {
	return ws.Conn.LocalAddr()
}

func (ws *WebsocketIO) RemoteAddr() net.Addr {
	return ws.Conn.RemoteAddr()
}

func (ws *WebsocketIO) SetDeadline(t time.Time) error {
	return ws.Conn.SetReadDeadline(t)
}

func (ws *WebsocketIO) SetReadDeadline(t time.Time) error {
	return ws.Conn.SetReadDeadline(t)
}

func (ws *WebsocketIO) SetWriteDeadline(t time.Time) error {
	return ws.Conn.SetWriteDeadline(t)
}
