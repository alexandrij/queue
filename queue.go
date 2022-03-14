package queue

import "github.com/alexandrij/linkedlist"

type Queue struct {
	data linkedlist.LinkedList
}

func (queue *Queue) Enqueue(v interface{}) {
	queue.data.Add(v)
}

func (queue *Queue) Dequeue() interface{} {
	if queue.Len() > 0 {
		return queue.data.RemoveHead().Value
	} else {
		return nil
	}
}

func (queue *Queue) Len() int {
	return queue.data.Len()
}

func New() *Queue {
	return new(Queue)
}
