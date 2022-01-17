package scheduler

import (
	"time"
)

type TaskOptions struct {
	Function         func() interface{}
	Frequency        time.Duration
	MaxRunInParallel int
	Name             string
}

type Task struct {
	f                func() interface{}
	frequency        time.Duration
	maxRunInParallel int
	name             string
}

func NewTask(options TaskOptions) *Task {
	return &Task{
		f:                options.Function,
		frequency:        options.Frequency,
		maxRunInParallel: options.MaxRunInParallel,
		name:             options.Name,
	}
}
