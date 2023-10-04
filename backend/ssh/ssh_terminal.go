package ssh

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/
#cgo CFLAGS: -I${SRCDIR}/../libs/
#cgo darwin LDFLAGS: -framework CoreFoundation -framework Security -lpthread -ldl -lm -lsectran-terminal
#cgo linux LDFLAGS: -lsectran-terminal -lpthread -ldl -lm
#include <stdlib.h>
#include "terminal.h"
*/
import "C"
import (
	"unsafe"

	"github.com/sirupsen/logrus"
)

// export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:/Users/ryan/Desktop/development/go/sectran/backend/terminal/lib/
// while true; do cat terminal.dump; sleep 1; clear; done
func XtermStart(width, heigth int) unsafe.Pointer {
	logrus.Infof("new terminal width:%d,height:%d", width, heigth)
	return unsafe.Pointer(C.sectran_terminal_alloc(C.int(width), C.int(heigth)))
}

func XtermGetCommand(terminal unsafe.Pointer) []byte {
	cchar := C.get_current_command((*C.sectran_terminal_handle)(unsafe.Pointer(terminal)))
	if cchar != nil {
		defer C.free(unsafe.Pointer(cchar))
		byteSlice := ([]byte)(C.GoString(cchar))
		return byteSlice
	}
	return nil
}

func XtermWrite(terminal unsafe.Pointer, data []byte) {
	C.sectran_terminal_write((*C.sectran_terminal_handle)(unsafe.Pointer(terminal)), (*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)))
}

func XtermFree(terminal unsafe.Pointer) {
	C.sectran_terminal_stop((*C.sectran_terminal_handle)(unsafe.Pointer(terminal)))
}
