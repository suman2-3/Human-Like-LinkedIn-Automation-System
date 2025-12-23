package connection

import (
	"time"
)

type Limiter struct {
	Max   int
	Count int
}

func NewLimiter(max int) *Limiter {
	return &Limiter{Max: max}
}

func (l *Limiter) Allow() bool {
	if l.Count >= l.Max {
		return false
	}
	l.Count++
	return true
}

func Cooldown() {
	time.Sleep(time.Duration(20+5*lrand()) * time.Second)
}

func lrand() int {
	return int(time.Now().UnixNano()%5 + 1)
}
