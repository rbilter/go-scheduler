package scheduler

type TaskScheduler struct {
	schedulers []*scheduler
	tasks      []*Task
}

func NewTaskScheduler(tasks []*Task) *TaskScheduler {
	return &TaskScheduler{
		schedulers: make([]*scheduler, len(tasks)),
		tasks:      tasks,
	}
}

func (s *TaskScheduler) Start() {
	for i, task := range s.tasks {
		s.schedulers[i] = newScheduler(schedulerOptions{
			f:                task.f,
			frequency:        task.frequency,
			maxRunInParallel: task.maxRunInParallel,
			name:             task.name,
		})
		s.schedulers[i].start()
	}
}

func (s *TaskScheduler) Stop() {
	for _, scheduler := range s.schedulers {
		scheduler.stop()
	}
}

func (s *TaskScheduler) TasksInProgress() (total int) {
	total = 0
	for _, scheduler := range s.schedulers {
		total += scheduler.tasksInProgress()
	}
	return total
}
