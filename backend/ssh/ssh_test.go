package ssh

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSSHClient(t *testing.T) {

}

func TestSSHServer(t *testing.T) {

}

// 启动测试（proxy）程序
func TestSSHProxy(t *testing.T) {
	config := SSHConfig{
		Port:            19527,
		Host:            "0.0.0.0",
		InteractiveAuth: true,
	}

	sm, err := StartSSHModule(&config)
	if err != nil {
		logrus.Infof("%+v\n", err)
	}

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	for {
		select {
		case res := <-sm.ResponseChan:
			if res.err != nil {
				logrus.Infof("recieve error info: %+v\n", res.err)
			}
		case <-signalChan:
			os.Exit(0)
		}
	}

}
