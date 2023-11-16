//go:build !ssh_commands_audit
// +build !ssh_commands_audit

package ssh

import "C"
import (
	"unsafe"
)

// This file lets us compile without including the real terminal
// implementation. Use the ssh_commands_audit build tag to include the
// real implementation.

// while true; do cat terminal.dump; sleep 1; clear; done
func XtermStart(width, heigth int) unsafe.Pointer {
	return unsafe.Pointer(nil)
}

func XtermGetCommand(terminal unsafe.Pointer) []byte {
	return nil
}

func XtermWrite(terminal unsafe.Pointer, data []byte) {

}

func XtermFree(terminal unsafe.Pointer) {

}

func XtermDumpToFile(terminal unsafe.Pointer) {

}

func XtermFlush(terminal unsafe.Pointer) {

}

func XtermMarkStdin(terminal unsafe.Pointer, data []byte) {

}
