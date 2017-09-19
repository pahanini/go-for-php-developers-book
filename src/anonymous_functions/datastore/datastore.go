package datastore

type Datastore struct {
	hitCounter int
}

func (ds *Datastore) RegisterHit() {
	ds.hitCounter++
}

func (ds *Datastore) TotalHits() int {
	return ds.hitCounter
}
