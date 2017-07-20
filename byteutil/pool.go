package byteutil

import (
	"sync"
	"bytes"
)

type BufferPool struct {
	pool sync.Pool
}

func NewBufferPool() *BufferPool {
	return &BufferPool{
		pool: sync.Pool{
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
