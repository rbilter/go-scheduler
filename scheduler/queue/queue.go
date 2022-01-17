package queue

import "github.com/google/uuid"

type QueueLength int

const (
	Infinite QueueLength = 0
)

type Queue struct {
	items []*uuid.UUID
	max   int
	op    chan bool
}

func NewQueue(max QueueLength) *Queue {
	q := make(chan bool)
	if max > 0 {
		q = make(chan bool, max)
	}
	return &Queue{
		max: int(max),
		op:  q,
	}
}

func (s *Queue) Length() int {
	return len(s.items)
}

func (s *Queue) Pop() *uuid.UUID {
	var item *uuid.UUID
	go func() {
		for {
			select {
			case <-s.op:
				if len(s.items) > 0 {
					item = s.items[0]
					s.items = s.items[1:]
				}
				return
			}
		}
	}()
	s.op <- true
	return item
}

func (s *Queue) Push(id *uuid.UUID) bool {
	pushed := false
	go func() {
		for {
			select {
			case <-s.op:
				if s.max == int(Infinite) {
					s.items = append(s.items, id)
					pushed = true
				} else {
					if len(s.items) < s.max {
						s.items = append(s.items, id)
						pushed = true
					}
				}
				return
			}
		}
	}()
	s.op <- true
	return pushed
}
