package main

import (
	"sync"
	"time"
)

type dataStore struct {
	sync.Mutex
	store map[string][]byte
}

func newDataStore() *dataStore {
	res := &dataStore{
		Mutex: sync.Mutex{},
		store: map[string][]byte{},
	}

	// clean store every hour
	go func(ds *dataStore) {
		ticker := time.NewTicker(time.Hour)
		for _ = range ticker.C {
			ds.forEach(func(k string, v []byte) {
				delete(ds.store, k)
			})
		}
	}(res)

	return res
}

func (ds *dataStore) get(key string) []byte {
	ds.Lock()
	defer ds.Unlock()

	res, ok := ds.store[key]
	if !ok {
		return nil
	}

	return res
}

func (ds *dataStore) set(key string, val []byte) {
	ds.Lock()
	defer ds.Unlock()
	ds.store[key] = val
}

func (ds *dataStore) delete(key string) bool {
	ds.Lock()
	defer ds.Unlock()

	if _, ok := ds.store[key]; !ok {
		return false
	}

	delete(ds.store, key)
	return true
}

func (ds *dataStore) forEach(f func(key string, val []byte)) {
	ds.Lock()
	defer ds.Unlock()

	for k, v := range ds.store {
		f(k, v)
	}
}
