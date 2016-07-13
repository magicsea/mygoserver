//author by xin.zhao

package pack

import (
	"errors"
)

/*
type:           [1]
iter:     begin(2)     end(8)
            |           |
data:   _ _ * * * * * * _ _ _
buffer: _ _ _ _ _ _ _ _ _ _ _
index:  0 1 2 3 4 5 6 7 8 9 10

type:           [2]
iter:      end(2)   begin(7)
            |         |
data:   * * _ _ _ _ _ * * * *
buffer: _ _ _ _ _ _ _ _ _ _ _
index:  0 1 2 3 4 5 6 7 8 9 10


type:           [3]
iter:      begin(4),end(4)
                |
data:   _ _ _ _ _ _ _ _ _ _ _
buffer: _ _ _ _ _ _ _ _ _ _ _
index:  0 1 2 3 4 5 6 7 8 9 10

type:           [4]
iter:      begin(4),end(4)
|				|
data:   * * * * * * * * * * *
buffer: _ _ _ _ _ _ _ _ _ _ _
index:  0 1 2 3 4 5 6 7 8 9 10

*/

type CircleBuffer struct {
	buffer         []byte
	datasize       int
	size           int
	begin          int
	end            int
	store_datasize int
	store_begin    int
	store_end      int
}

func (b *CircleBuffer) Ini(n int) bool {
	b.Reset(n)
	return true
}

func (b *CircleBuffer) Reset(n int) {
	b.datasize = 0
	b.size = n
	b.begin = 0
	b.end = 0
	b.store_datasize = 0
	b.store_begin = 0
	b.store_end = 0
	b.buffer = make([]byte, n)
}

func (b *CircleBuffer) CanWrite(size int) bool {
	return b.datasize+size <= b.size
}

func (b *CircleBuffer) SkipWrite(size int) {
	b.datasize += size
	b.end += size
	if b.end >= b.size {
		b.end -= b.size
	}
}

func (b *CircleBuffer) Write(data []byte) (n int, err error) {

	size := len(data)

	if !b.CanWrite(size) {

		n = 0
		err = errors.New(" buf is flow ...")
		return
	}

	b.realWrite(data)

	b.SkipWrite(size)

	n = size
	err = nil
	return
}

func (b *CircleBuffer) CanRead(size int) bool {
	return b.datasize >= size
}

func (b *CircleBuffer) SkipRead(size int) {
	b.datasize -= size
	b.begin += size
	if b.begin >= b.size {
		b.begin -= b.size
	}
}

func (b *CircleBuffer) Read(data []byte) (n int, err error) {

	size := len(data)

	if !b.CanRead(size) {

		n = 0
		err = errors.New("not enough read byte")

		return
	}

	b.realRead(data)

	b.SkipRead(size)

	n = size
	err = nil
	return
}

func (b *CircleBuffer) Store() {
	b.store_datasize = b.datasize
	b.store_begin = b.begin
	b.store_end = b.end
}

func (b *CircleBuffer) Restore() {
	b.datasize = b.store_datasize
	b.begin = b.store_begin
	b.end = b.store_end
}

func (b *CircleBuffer) Clear() {
	b.datasize = 0
	b.begin = 0
	b.end = 0
}

func (b *CircleBuffer) Size() int {
	return b.datasize
}

func (b *CircleBuffer) Capacity() int {
	return b.size
}

func (b *CircleBuffer) Empty() bool {
	return b.Size() == 0
}

func (b *CircleBuffer) Full() bool {
	return b.Size() == b.Capacity()
}

func (b *CircleBuffer) FreeSpace() int {
	return b.Capacity() - b.Size()
}

func (b *CircleBuffer) realWrite(data []byte) {

	size := len(data)

	if b.end >= b.begin {
		// [1][3]
		// 能装下
		if b.size-b.end >= size {
			copy(b.buffer[b.end:], data)
		} else {
			copy(b.buffer[b.end:], data)
			copy(b.buffer, data[(b.size-b.end):])
		}
	} else {
		//[2]
		copy(b.buffer[b.end:], data)
	}
}

func (b *CircleBuffer) realRead(data []byte) {

	size := len(data)

	if b.begin >= b.end {
		// [2][4]
		// 能读完
		if b.size-b.begin >= size {
			copy(data, b.buffer[b.begin:])
		} else {
			copy(data, b.buffer[b.begin:])
			copy(data[(b.size-b.begin):], b.buffer[0:])
		}
	} else {
		// [1]
		copy(data, b.buffer[b.begin:])
	}
}

func (b *CircleBuffer) Output(outfunc func([]byte) (int, error)) (n int, err error) {

	if b.Empty() {
		return 0, nil
	}

	if b.begin >= b.end {
		// [2][4]
		n, err = outfunc(b.buffer[b.begin:b.size])
		if err != nil {
			return
		}

		if n == b.size-b.begin {
			n1, err1 := outfunc(b.buffer[0:b.end])
			n += n1
			err = err1
			if err != nil {
				return
			}
		}
	} else {
		// [1]
		n, err = outfunc(b.buffer[b.begin:b.end])
	}

	b.SkipRead(n)

	return
}

func (b *CircleBuffer) Input(infunc func([]byte) (int, error)) (n int, err error) {

	if b.Full() {
		return 0, nil
	}

	if b.end >= b.begin {
		// [1][3]
		n, err = infunc(b.buffer[b.end:b.size])
		if err != nil {
			return
		}

		if n == b.size-b.end {
			n1, err1 := infunc(b.buffer[0:b.begin])
			n += n1
			err = err1
			if err != nil {
				return
			}
		}
	} else {
		//[2]
		n, err = infunc(b.buffer[b.end:b.begin])
	}

	b.SkipWrite(n)

	return
}
