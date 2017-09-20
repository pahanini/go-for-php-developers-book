package datastore

import (
	"sync"
)

type Datastore struct {
	hitCounter int
	mu         sync.RWMutex
}

func (ds *Datastore) RegisterHit() {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.hitCounter++
}

func (ds *Datastore) TotalHits() int {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.hitCounter
}
