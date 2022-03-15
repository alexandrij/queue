package sync

import (
	"testing"
)

func checkLen(t *testing.T, queue *Queue, len int) bool {
	if n := queue.Len(); n != len {
		t.Errorf("queue.Len() = %v, want %v", n, len)
		return false
	}
	return true
}

func TestQueue_Len(t *testing.T) {
	queue := New()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	checkLen(t, queue, 3)

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	checkLen(t, queue, 0)

	queue.Dequeue()
	checkLen(t, queue, 0)
}

func TestQueue_Dequeue(t *testing.T) {
	q := New()
	size := 10

	for i := 0; i < size; i++ {
		q.Enqueue(i)
	}

	for i, v := 0, q.Dequeue(); v != nil; v = q.Dequeue() {
		if v != i {
			t.Errorf("queue.Dequeue() = %#v, want %#v", v, i)
		}
		i++
	}

	end := q.Dequeue()
	if end != nil {
		t.Errorf("queue.Dequeue() = %#v, want %#v", end, nil)
	}

	checkLen(t, q, 0)
}
