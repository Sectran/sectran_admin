//go:build terminal
// +build terminal

package ssh

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/
#cgo CFLAGS: -I${SRCDIR}/../libs/
#cgo darwin LDFLAGS: -framework CoreFoundation -framework Security -lpthread -ldl -lm -lsectran-terminal
#cgo linux LDFLAGS: -lsectran-terminal -lpthread -ldl -lm
#include <stdlib.h>
#include "terminal.h"

#ifdef DEBUG
void sectran_terminal_print_to_file(const sectran_terminal* terminal)
{
    const char* filename = "terminal.dump";
    if (!filename || !terminal || !terminal->buffer)
    {
        printf("Invalid arguments or uninitialized terminal.\n");
        return;
    }

    FILE* file = fopen(filename, "w");
    if (file == NULL)
    {
        printf("Unable to open file %s for writing.\n", filename);
        return;
    }

    printf("terminal cursor_row:%d\n",terminal->cursor_row);

    for (int j = 0; j <= terminal->cursor_row; j++)
    {
        sectran_buffer_row row = terminal->buffer->rows[j];
        fprintf(file, "%d:", j);
        for (int i = 0; i < row.length; i++)
        {
            fprintf(file, "%c", row.chars[i]);
        }

        fprintf(file, "\n");
    }

    fflush(stdout);
    fclose(file);
}
#endif
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
	cchar := C.get_current_command((terminal))
	if cchar != nil {
		defer C.free(unsafe.Pointer(cchar))
		byteSlice := ([]byte)(C.GoString(cchar))
		return byteSlice
	}
	return nil
}

func XtermWrite(termianl unsafe.Pointer, data []byte) {
	C.sectran_terminal_write((termianl), (*C.char)(unsafe.Pointer(&data[0])), C.int(len(data)))
	// C.sectran_terminal_print_to_file((termianl))
}

func XtermFree(termianl unsafe.Pointer) {
	C.sectran_terminal_stop(termianl)
}
