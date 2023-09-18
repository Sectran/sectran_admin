package signal

import (
	"os"
	"os/signal"
	"syscall"
)

func ReigisterSignal() chan os.Signal {
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	return signalChan
}
