// Based upon sync.WaitGroup, SizedWaitGroup allows to start multiple
// routines and to wait for their end using the simple API.

// SizedWaitGroup adds the feature of limiting the maximum number of
// concurrently started routines. It could for example be used to start
// multiples routines querying a database but without sending too much
// queries in order to not overload the given database.
//
// Rémy Mathieu © 2016
package sizedwaitgroup

import (
	"sync"
)

// SizedWaitGroup has the same role and close to the
// same API as the Golang sync.WaitGroup but adds a limit of
// the amount of goroutines started concurrently.
type SizedWaitGroup struct {
	Size int

	current chan bool
	wg      sync.WaitGroup
}

// New creates a SizedWaitGroup.
// The limit parameter if the maximum amount of
// goroutine which can be started concurrently.
func New(limit int) SizedWaitGroup {
	size := 4294967295 // 2^32 - 1
	if limit > 0 {
		size = limit
	}
	return SizedWaitGroup{
		Size: size,

		current: make(chan bool, size),
		wg:      sync.WaitGroup{},
	}
}

// Add increments the internal WaitGroup counter.
// It can be blocking if the limit of spawned goroutine
// has been reached. It will stop blocking when Done has
// been called.
//
// See sync.WaitGroup documentation for more information.
func (s *SizedWaitGroup) Add() {
	s.current <- true
	s.wg.Add(1)
}

// Done decrements the SizedWaitGroup counter.
// See sync.WaitGroup documentation for more information.
func (s *SizedWaitGroup) Done() {
	<-s.current
	s.wg.Done()
}

// Wait blocks until the SizedWaitGroup counter is zero.
// See sync.WaitGroup documentation for more information.
func (s *SizedWaitGroup) Wait() {
	s.wg.Wait()
}
