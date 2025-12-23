package stealth

import (
	"math/rand"
	"time"
)

var actionCount = 0

// PauseBetweenActions simulates human pauses
func PauseBetweenActions() {
	actionCount++

	// Short pause after every action
	time.Sleep(time.Duration(rand.Intn(1500)+800) * time.Millisecond)

	// Medium break every 5 actions
	if actionCount%5 == 0 {
		time.Sleep(time.Duration(rand.Intn(6000)+4000) * time.Millisecond)
	}

	// Long break every 15 actions
	if actionCount%15 == 0 {
		time.Sleep(time.Duration(rand.Intn(15000)+10000) * time.Millisecond)
	}
}
