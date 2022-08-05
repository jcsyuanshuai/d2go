package mem

import "C"
import "unsafe"

type Buf struct {
	Next *Buf
	Cap  int
	len  int
	head int
	data unsafe.Pointer
}

func NewBuf(size int) *Buf {
	return &Buf{
		Cap:  size,
		len:  0,
		head: 0,
		Next: nil,
		data: Malloc(size),
	}
}

func (b *Buf) SetBytes(src []byte) {
	Memcpy(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), src, len(src))
	b.len += len(src)
}

func (b *Buf) GetBytes() (data []byte) {
	data = C.GoBytes(unsafe.Pointer(uintptr(b.data)+uintptr(b.head)), C.int(b.len))
	return
}

func (b *Buf) Copy(other *Buf) {
	Memcpy(b.data, other.GetBytes(), other.len)
	b.head = 0
	b.len = other.len
}
