package conn

import (
	"net"
	"net/http"

	"github.com/sirupsen/logrus"
)

type serve func(listener net.Listener, connChan chan net.Conn)

func AcceptConnection(l net.Listener, websocket bool) <-chan net.Conn {
	var ch chan net.Conn = make(chan net.Conn, 1)
	var acceptServe serve

	if websocket {
		acceptServe = func(l net.Listener, ch chan net.Conn) {
			for {
				http.Serve(l, &WebScoketServer{
					Chan:     ch,
					Listener: l,
				})
			}
		}
	} else {
		acceptServe = func(l net.Listener, ch chan net.Conn) {
			for {
				c, err := l.Accept()
				if err != nil {
					logrus.Errorf("error accept client due to %s", err)
					return
				}
				ch <- c
			}
		}
	}

	go acceptServe(l, ch)

	return ch
}
