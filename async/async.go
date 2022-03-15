package async

import (
	squeue "queue/sync"
	"sync"
)

type Queue struct {
	squeue.Queue
	mx sync.Mutex
}

func (queue *Queue) Enqueue(value interface{}) {
	queue.mx.Lock()
	queue.Queue.Enqueue(value)
	queue.mx.Unlock()
}

func (queue *Queue) Dequeue() interface{} {
	queue.mx.Lock()
	value := queue.Queue.Dequeue()
	queue.mx.Unlock()
	return value
}

func New() *Queue {
	return new(Queue)
}
