package websocket

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/gravitational/trace"
)

type WebsocketIO struct {
	*websocket.Conn
	remaining []byte
}

func (ws *WebsocketIO) Write(p []byte) (int, error) {
	err := ws.WriteMessage(websocket.BinaryMessage, p)
	if err != nil {
		return 0, trace.Wrap(err)
	}

	return len(p), nil
}

func (ws *WebsocketIO) Read(p []byte) (int, error) {
	if len(ws.remaining) == 0 {
		ty, data, err := ws.ReadMessage()
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

func (ws *WebsocketIO) SetDeadline(t time.Time) error {
	return ws.Conn.SetReadDeadline(t)
}
