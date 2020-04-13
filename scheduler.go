package discordbot

import "time"

type Scheduler struct {
	TimeZone string
	Tasks    []SchedulerTask
}

type SchedulerTask struct {
	ID   string
	Time time.Time
	Task func()
}

func NewScheduler() Scheduler {
	return Scheduler{}
}

func (s *Scheduler) List() error {
	return nil
}

func (s *Scheduler) Add(f func()) error {
	for _, task := range s.Tasks {
		go task.Task()
	}
	return nil
}

func (s *Scheduler) Fix(id string) error {
	return nil
}

func (s *Scheduler) Remove() error {
	return nil
}

func (s *Scheduler) Run() error {
	return nil
}
