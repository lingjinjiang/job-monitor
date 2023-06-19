package message

type localQueue struct {
	cache chan Event
}

func (q *localQueue) Push(e Event) error {
	q.cache <- e
	return nil
}

func (q *localQueue) Pop() (Event, error) {
	return <-q.cache, nil
}
