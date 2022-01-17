package scheduler

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rbilter/go-scheduler/scheduler/queue"
)

type schedulerOptions struct {
	f                func() interface{}
	frequency        time.Duration
	maxRunInParallel int
	name             string
}

type scheduler struct {
	done             chan bool
	f                func() interface{}
	frequency        time.Duration
	inProgress       *queue.Queue
	maxRunInParallel int
	name             string
	startTime        time.Time
}

func newScheduler(options schedulerOptions) *scheduler {
	return &scheduler{
		done:             make(chan bool),
		f:                options.f,
		frequency:        options.frequency,
		inProgress:       queue.NewQueue(queue.Infinite),
		maxRunInParallel: options.maxRunInParallel,
		name:             options.name,
	}
}

func (s *scheduler) start() []interface{} {
	tick := time.NewTicker(s.frequency)
	response := make([]interface{}, 0)
	go func() {
		for {
			select {
			case <-s.done:
				tick.Stop()
				return
			case <-tick.C:
				go func() {
					uuid := uuid.New()
					if s.inProgress.Push(&uuid) {
						fmt.Printf("Scheduler '%s': task {%s} started\n", s.name, uuid)
						response = append(response, s.f())
						s.inProgress.Pop()
						fmt.Printf("Scheduler '%s': task {%s} ended\n", s.name, uuid)
					}
				}()
			}
		}
	}()
	return response
}

func (s *scheduler) stop() {
	s.done <- true
}

func (s *scheduler) tasksInProgress() int {
	return s.inProgress.Length()
}
