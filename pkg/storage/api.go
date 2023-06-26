package storage

import (
	"job-monitor/pkg/api"
	"job-monitor/pkg/message"
	"job-monitor/pkg/storage/local"
)

type Storage interface {
	AddJob(api.Job) error
	UpdateJob(api.Job) error
	DeleteJob(api.Job) error
	LoadJobs() []api.Job
}

func NewStorage(mode string, queue message.Queue) Storage {
	s := local.NewLocalFileStorage(queue, "/tmp/data")
	return s
}
