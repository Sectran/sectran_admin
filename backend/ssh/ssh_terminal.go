package ssh

/*
#cgo LDFLAGS: -L${SRCDIR}/../terminal/libs/
#cgo CFLAGS: -I${SRCDIR}/../terminal/libs/
#cgo darwin LDFLAGS: -framework CoreFoundation -framework Security -lpthread -ldl -lm -lsectran-terminal
#cgo linux LDFLAGS: -lsectran-terminal -lpthread -ldl -lm
#include <stdlib.h>

typedef void *sectran_terminal_handle;
sectran_terminal_handle *sectran_terminal_alloc(int width, int height);
int sectran_terminal_write(sectran_terminal_handle *terminal, const char *c,
                           int size);
void sectran_terminal_stop(sectran_terminal_handle *terminal);
char *get_current_command(sectran_terminal_handle *term);
void sectran_terminal_print_to_file(sectran_terminal_handle* terminal);

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

func DumpToFile(terminal unsafe.Pointer) {
	C.sectran_terminal_print_to_file((*C.sectran_terminal_handle)(unsafe.Pointer(terminal)))
}
