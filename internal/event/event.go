package event

import (
	"log/slog"
	"slices"
	"sync"
)

type Distribution interface {
	Send(info string)
	Register(ch chan string)
	Unregister(ch chan string)
}

type Distributer struct {
	logger    *slog.Logger
	mu        sync.Mutex
	destChans []chan string
}

func NewDistributer(logger *slog.Logger) (Distribution, error) {
	ed := &Distributer{
		logger:    logger,
		mu:        sync.Mutex{},
		destChans: make([]chan string, 0),
	}
	return ed, nil
}

func (ed *Distributer) Send(payload string) {
	ed.logger.Debug(
		"sending events to channels",
		"ch_count", len(ed.destChans),
	)

	ed.mu.Lock()
	defer ed.mu.Unlock()

	for _, destCh := range ed.destChans {
		go func(ch chan string) {
			ch <- payload
		}(destCh)
	}
}

func (ed *Distributer) Register(ch chan string) {
	ed.logger.Info(
		"registering channel",
		"ch_count_before", len(ed.destChans),
		"ch_count_after", len(ed.destChans)+1,
	)

	ed.mu.Lock()
	defer ed.mu.Unlock()

	ed.destChans = append(ed.destChans, ch)
}

func (ed *Distributer) Unregister(ch chan string) {
	ed.logger.Info(
		"unregistering channel",
		"ch_count_before", len(ed.destChans),
		"ch_count_after", len(ed.destChans)-1,
	)

	ed.mu.Lock()
	defer ed.mu.Unlock()

	for i, destCh := range ed.destChans {
		if ch == destCh {
			ed.destChans = slices.Delete(ed.destChans, i, i+1)
			break
		}
	}
}
