package store

import "time"

type Store interface {
	Set(key []byte, value []byte) (err error)
	Get(key []byte) (value []byte, err error)
	GetAll(prefix []byte) (keyValues map[string][]byte, err error)
	Count(prefix []byte) (count uint, err error)
	Delete(key []byte) (err error)
	Close() (err error)
	Stats() Stats
}

type Stats struct {
	Hits           uint64
	Misses         uint64
	Evictions      uint64
	EvictionWeight uint64
	LoadSuccesses  uint64
	LoadFailures   uint64
	TotalLoadTime  time.Duration
}
