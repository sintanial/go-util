package syncutil

import (
	"time"
	"sync"
)

type PeriodicalOnce struct {
	Period time.Duration

	mu        sync.Mutex
	executeAt time.Time
}

func (self *PeriodicalOnce) Do(fn func()) {
	self.mu.Lock()
	if time.Now().Sub(self.executeAt) < self.Period {
		self.mu.Unlock()
		return
	}

	self.executeAt = time.Now()
	self.mu.Unlock()

	fn()
}
