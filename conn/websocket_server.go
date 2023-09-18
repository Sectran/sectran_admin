package conn

import (
	"net"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WebScoketServer struct {
	Chan     chan net.Conn
	Listener net.Listener
}

func (ws *WebScoketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upgrade := &websocket.Upgrader{
		ReadBufferSize:  12288,
		WriteBufferSize: 12288,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("websocket error:%s", err)
		return
	}

	webConn := &WebsocketIO{}
	webConn.Conn = conn

	ws.Chan <- webConn
}
