package ecs

import "container/list"

// Queue is a queue
type Queue interface {
	Front() *list.Element
	Len() int
	Add(any)
	Remove()
}

type queueImpl struct {
	*list.List
}

func (q *queueImpl) Add(v any) {
	q.PushBack(v)
}

func (q *queueImpl) Remove() {
	e := q.Front()
	q.List.Remove(e)
}

// New is a new instance of a Queue
func NewQueue() Queue {
	return &queueImpl{list.New()}
}
