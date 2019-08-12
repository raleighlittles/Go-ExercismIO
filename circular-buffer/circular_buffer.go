package circular

import (
	"errors"
)

type Buffer struct {
	 arr []byte
	 read_index int
	 write_index int
	 num_elements int
}

func NewBuffer(size int) *Buffer {
	buf := Buffer{make([]byte, size), 0, 0, 0}
	return &buf
}

func (b *Buffer) ReadByte() (result byte, err error) {
	if (len(b.arr) == 0) || (b.num_elements == 0) {
		return 0, errors.New("Can't read out of empty buffer")
	}

	result = b.arr[b.read_index]
	err = nil

	b.read_index = (b.read_index + 1) % len(b.arr)
	b.num_elements--

	return

}

func (b *Buffer) WriteByte(c byte) (error) {
	if (len(b.arr) == 0) {
		return errors.New("Can't write to an empty buffer")
	}

	if (len(b.arr) == b.num_elements) {
		return errors.New("Can't write to full buffer")
	}

	b.arr[b.write_index] = c
	b.write_index = (b.write_index + 1) % len(b.arr)
	b.num_elements++

	return nil

}

func (b *Buffer) Overwrite(c byte) {

	if (len(b.arr) == b.num_elements) {
		// array is full. overwrite replaces
		b.arr[b.read_index] = c
		b.read_index = (b.read_index + 1) % len(b.arr)
		return
	}

	b.arr[b.write_index] = c

	b.write_index = (b.write_index + 1) % len(b.arr)

	b.num_elements++

	return

}

func (b *Buffer) Reset() {
	cur_length := len(b.arr)
	b.arr = nil
	b.arr = make([]byte, cur_length)
	b.read_index = 0
	b.write_index = 0
	b.num_elements = 0
}
