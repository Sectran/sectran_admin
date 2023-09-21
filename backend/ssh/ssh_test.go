package ssh

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSSHClient(t *testing.T) {

}

func TestSSHServer(t *testing.T) {

}

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

	res := <-sm.ResponseChan
	if res.err != nil {
		logrus.Infof("%+v\n", res.err)
	}

	select {}
}
