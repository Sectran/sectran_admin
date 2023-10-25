package main

import (
	"os"
	"os/signal"
	"sectran/backend/ssh"
	"syscall"

	"github.com/sirupsen/logrus"
)

// encoded[i] = arr[i] XOR arr[i + 1]
func Decode_xor_array(encode_array []uint32, origin_first_item uint32) []uint32 {
	origin_array := make([]uint32, len(encode_array)+1)
	origin_array[0] = origin_first_item

	for i, v := range encode_array {
		origin_array[i+1] = v ^ origin_array[i]
	}
	return origin_array
}

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
	//fmt.Println(user)
}
