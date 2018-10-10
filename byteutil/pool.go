package byteutil

import (
	"sync"
	"bytes"
	"bufio"
)

type BufferPool struct {
	pool *sync.Pool
}

func NewBufferPool() *BufferPool {
	return &BufferPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

func (self *BufferPool) Get() *bytes.Buffer {
	return self.pool.Get().(*bytes.Buffer)
}

func (self *BufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	self.pool.Put(b)
}

type BufioReaderPool struct {
	pool *sync.Pool
}

func NewBufioReaderPool(size int) *BufioReaderPool {
	return &BufioReaderPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return bufio.NewReaderSize(nil, size)
			},
		},
	}
}

func (self *BufioReaderPool) Get() *bufio.Reader {
	return self.pool.Get().(*bufio.Reader)
}

func (self *BufioReaderPool) Put(b *bufio.Reader) {
	b.Reset(nil)
	self.pool.Put(b)
}


type BytePool struct {
	pool *sync.Pool
}

func NewBytePool(size int) *BytePool {
	return &BytePool{
		pool: &sync.Pool{
			New: func() interface{} {
				return make([]byte, size)
			},
		},
	}
}

func (self *BytePool) Get() []byte {
	return self.pool.Get().([]byte)
}

func (self *BytePool) Put(b []byte) {
	self.pool.Put(b)
}