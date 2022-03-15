package async

import (
	"sync"
	"testing"
)

func TestAsyncQueue_Enqueue(t *testing.T) {
	wg := sync.WaitGroup{}

	asyncQueue := New()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			asyncQueue.Enqueue(i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	if n := asyncQueue.Len(); n != 10 {
		t.Errorf("queue.Len() = %v, want %v", n, 10)
	}
}
