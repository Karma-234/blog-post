package main

import "sync"

type safeCounter struct {
	counts map[string]int
	mux    sync.Mutex
	rmux   sync.RWMutex // Read write mutex allows infinite reads of a data but not multiple writes
}

func (r *safeCounter) slowIncrement(key string) {}
func (r *safeCounter) inc(key string) {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.slowIncrement(key)

}
func (r *safeCounter) returnCount(key string) int {
	r.rmux.RLock()
	defer r.rmux.Unlock()
	return r.counts[key]
}
func race() {

}

// Mutex is used to safely access data concurrently. This is used to prevent multiple go routines from accessing and writing data at the same time.
// Genrics. You get the idea

// Any is an alis for interface
func generic[T interface{}, C any](s []T) ([]T, []T) {
	mid := len(s) / 2

	return s[:mid], s[mid:]
}
