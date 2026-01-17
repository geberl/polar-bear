package store

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/maypok86/otter/v2"
	"github.com/maypok86/otter/v2/stats"
)

type OtterStore struct {
	logger  *slog.Logger
	cache   *otter.Cache[string, string]
	counter *stats.Counter
}

func NewOtterStore() (Store, error) {
	counter := stats.NewCounter()

	o, err := otter.New(&otter.Options[string, string]{
		MaximumSize:   10_000,
		StatsRecorder: counter,
	})

	if err != nil {
		return nil, err
	}

	os := &OtterStore{
		logger:  slog.With("component", "otter-store"),
		cache:   o,
		counter: counter,
	}

	return os, nil
}

func (os *OtterStore) Set(key []byte, value []byte) error {
	os.cache.Set(string(key), string(value))
	return nil
}

func (os *OtterStore) Get(key []byte) ([]byte, error) {
	if value, ok := os.cache.GetIfPresent(string(key)); ok {
		return []byte(value), nil
	}
	return nil, fmt.Errorf("not found")
}

func (os *OtterStore) GetAll(prefix []byte) (map[string][]byte, error) {
	results := make(map[string][]byte)

	for key, value := range os.cache.All() {
		if strings.HasPrefix(key, string(prefix)) {
			results[string(key)] = []byte(value)
		}
	}

	return results, nil
}

func (os *OtterStore) Count(prefix []byte) (uint, error) {
	count := uint(0)

	for key := range os.cache.All() {
		if strings.HasPrefix(key, string(prefix)) {
			count++
		}
	}

	return count, nil
}

func (os *OtterStore) Delete(key []byte) error {
	if _, invalidated := os.cache.Invalidate(string(key)); invalidated {
		return nil
	}
	return fmt.Errorf("not deleted")
}

func (os *OtterStore) Close() error {
	return nil
}

func (os *OtterStore) Stats() Stats {
	snapshot := os.counter.Snapshot()
	return Stats{
		Hits:           snapshot.Hits,
		Misses:         snapshot.Misses,
		Evictions:      snapshot.Evictions,
		EvictionWeight: snapshot.EvictionWeight,
		LoadSuccesses:  snapshot.LoadSuccesses,
		LoadFailures:   snapshot.LoadFailures,
		TotalLoadTime:  snapshot.TotalLoadTime,
	}
}
