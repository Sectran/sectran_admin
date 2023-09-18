package utils

import (
	"bytes"
	"encoding/binary"
)

type Reader struct {
	buf *bytes.Reader
}

func NewReader(data []byte) *Reader {
	return &Reader{buf: bytes.NewReader(data)}
}

func (r *Reader) RemainLength() int {
	return r.buf.Len()
}

func (r *Reader) CheckLength(len int) int {
	var unread int = r.buf.Len()
	if unread < len {
		return -1
	}
	return unread
}

func (r *Reader) ReadBytes(length int) ([]byte, error) {
	buf := make([]byte, length)
	n, err := r.buf.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func (r *Reader) ReadByte() (byte, error) {
	n, err := r.buf.ReadByte()
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (r *Reader) ReadBigEndian16() (uint16, error) {
	var value uint16
	err := binary.Read(r.buf, binary.BigEndian, &value)
	return value, err
}

func (r *Reader) ReadBigEndian32() (uint32, error) {
	var value uint32
	err := binary.Read(r.buf, binary.BigEndian, &value)
	return value, err
}

func (r *Reader) ReadBigEndian64() (uint64, error) {
	var value uint64
	err := binary.Read(r.buf, binary.BigEndian, &value)
	return value, err
}

func (r *Reader) ReadLittleEndian16() (uint16, error) {
	var value uint16
	err := binary.Read(r.buf, binary.LittleEndian, &value)
	return value, err
}

func (r *Reader) ReadLittleEndian32() (uint32, error) {
	var value uint32
	err := binary.Read(r.buf, binary.LittleEndian, &value)
	return value, err
}

func (r *Reader) ReadLittleEndian64() (uint64, error) {
	var value uint64
	err := binary.Read(r.buf, binary.LittleEndian, &value)
	return value, err
}
