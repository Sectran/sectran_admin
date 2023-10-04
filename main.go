package main

import (
	"os"
	"os/signal"
	"sectran/backend/ssh"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	config := ssh.SSHConfig{
		Port:            19527,
		Host:            "0.0.0.0",
		InteractiveAuth: true,
	}

	sm, err := ssh.StartSSHModule(&config)
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
			if res.Err != nil {
				logrus.Infof("recieve error info: %+v\n", res.Err)
			}
		case <-signalChan:
			os.Exit(0)
		}
	}
}
