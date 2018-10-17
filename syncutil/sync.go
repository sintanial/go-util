package syncutil

import (
	"time"
	"sync"
)

type PeriodicalOnce struct {
	mu        sync.Mutex
	executeAt time.Time
}

func (self *PeriodicalOnce) Do(period time.Duration, fn func()) {
	self.mu.Lock()
	if time.Now().Sub(self.executeAt) < period {
		self.mu.Unlock()
		return
	}

	self.executeAt = time.Now()
	self.mu.Unlock()

	fn()
}
