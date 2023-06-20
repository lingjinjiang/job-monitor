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
}

func NewStorage(mode string, queue message.Queue) Storage {
	s := local.LoadFileStorage{Queue: queue}
	message.Register("add", s.AddJob)
	message.Register("update", s.UpdateJob)
	message.Register("delete", s.DeleteJob)
	return &s
}
