package ssh

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestSSHClient(t *testing.T) {
	buffer := &bytes.Buffer{}
	writer := bufio.NewWriter(buffer)

	terminal := "xterm256"
	// terminal length
	binary.Write(writer, binary.BigEndian, uint32(len(terminal)))
	// terminal string
	writer.WriteString(terminal)
	writer.Flush()

	fmt.Printf("b: %v\n", buffer.Bytes())
}
