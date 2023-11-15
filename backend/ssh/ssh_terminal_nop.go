//go:build !ssh_commands_audit
// +build !ssh_commands_audit

package ssh

import "C"
import (
	"unsafe"

	"github.com/sirupsen/logrus"
)

// This file lets us compile without including the real terminal
// implementation. Use the ssh_commands_audit build tag to include the
// real implementation.

// while true; do cat terminal.dump; sleep 1; clear; done
func XtermStart(width, heigth int) unsafe.Pointer {
	logrus.Errorln("build without ssh command audit module")
	return unsafe.Pointer(nil)
}

func XtermGetCommand(terminal unsafe.Pointer) []byte {
	logrus.Errorln("build without ssh command audit module")
	return nil
}

func XtermWrite(terminal unsafe.Pointer, data []byte) {
	logrus.Errorln("build without ssh command audit module")
}

func XtermFree(terminal unsafe.Pointer) {
	logrus.Errorln("build without ssh command audit module")
}

func XtermDumpToFile(terminal unsafe.Pointer) {
	logrus.Errorln("build without ssh command audit module")
}

func XtermFlush(terminal unsafe.Pointer) {
	logrus.Errorln("build without ssh command audit module")
}

func XtermMarkStdin(terminal unsafe.Pointer, data []byte) {
	logrus.Errorln("build without ssh command audit module")
}
