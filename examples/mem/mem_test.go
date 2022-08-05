package mem

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
	"unsafe"
)

func IsLittleEndian() bool {
	n := 0x01020304
	u := unsafe.Pointer(&n)
	ret := *(*byte)(u)
	return ret == 0x04
}

func IntToBytes(n uint32) []byte {
	x := int32(n)
	buffer := bytes.NewBuffer([]byte{})

	var order binary.ByteOrder
	if IsLittleEndian() {
		order = binary.LittleEndian
	} else {
		order = binary.BigEndian
	}
	binary.Write(buffer, order, x)
	return buffer.Bytes()
}

func TestMem(t *testing.T) {
	data := Malloc(4)
	fmt.Println("memory malloc 4 byte (32 bit)")
	myData := (*uint32)(data)
	fmt.Printf("memory value=%+3v,type=%T.\n", *myData, *myData)
	*myData = 5
	fmt.Printf("memory value=%+3v,type=%T.\n", *myData, *myData)
	x := uint32(100)
	Memcpy(data, IntToBytes(x), 4)
	fmt.Printf("memory value=%+3v,type=%T.\n", *myData, *myData)
	Free(data)
}
