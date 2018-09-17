package loadbalancer

import (
	"strconv"
	"sync"
	"testing"
)

func TestQueue(t *testing.T) {
	mutex := sync.RWMutex{}
	q := NewQueue(&mutex)

	for i := 0; i < 9; i++ {
		q.PushBack(&Request{
			data: "test" + strconv.Itoa(i),
		})
	}

	r := q.PopFront()
	if r.(*Request).data != "test0" {
		t.Errorf("Data should be 'test0', instead of %s", r.(*Request).data)
	}
}
