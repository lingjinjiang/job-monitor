package spark

import (
	"job-monitor/pkg/message"
)

type eventHandler struct {
	Queue message.Queue
}

func (h *eventHandler) addSparkApplication(obj interface{}) {
	app := obj.(*SparkApplication)
	h.Queue.Push(message.Event{Operation: "add", Job: app.Convert()})
}

func (h *eventHandler) updateSparkApplication(oldObj interface{}, newObj interface{}) {
	app := newObj.(*SparkApplication)
	h.Queue.Push(message.Event{Operation: "update", Job: app.Convert()})
}

func (h *eventHandler) deleteSparkApplication(obj interface{}) {
	app := obj.(*SparkApplication)
	h.Queue.Push(message.Event{Operation: "delete", Job: app.Convert()})
}
