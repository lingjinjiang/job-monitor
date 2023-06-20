package local

import (
	"job-monitor/pkg/api"
	"job-monitor/pkg/message"
)

type LoadFileStorage struct {
	Queue message.Queue
}

func (s *LoadFileStorage) AddJob(job api.Job) error {
	return nil
}

func (s *LoadFileStorage) UpdateJob(job api.Job) error {
	return nil
}

func (s *LoadFileStorage) DeleteJob(job api.Job) error {
	return nil
}
