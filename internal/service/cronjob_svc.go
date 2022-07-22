package service

import (
	"github.com/hibiken/asynq"

	"github.com/gole-dev/gole-layout/internal/tasks"
)

var DefaultJobs map[string]*asynq.Task

type JobFunc func()

type CronJobService struct {
}

func NewCronJobService() *CronJobService {
	return &CronJobService{}
}

// RegisterTask register task
func (s *CronJobService) RegisterTask() {
	DefaultJobs = map[string]*asynq.Task{
		tasks.TypeEmailWelcome: tasks.NewEmailWelcomeTask(1),
	}
}
