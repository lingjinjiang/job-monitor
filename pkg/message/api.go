package message

import "job-monitor/pkg/api"

type Event struct {
	Operation string
	Job       api.Job
}

type Queue interface {
	Push(Event) error
	Pop() (Event, error)
}

type eventHandler func(api.Job) error

var handlers map[string]eventHandler = make(map[string]eventHandler)

func Register(operation string, handler eventHandler) {
	handlers[operation] = handler
}

func NewQueue(mode string) Queue {
	return &localQueue{
		cache: make(chan Event, 10),
	}
}
