package loadbalancer

import (
	"fmt"
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
	fmt.Println(q.repo)

	for i := 0; i < 9; i++ {
		q.PopBack()
	}

	fmt.Println(q.repo)

	for i := 0; i < 9; i++ {
		q.PushBack(&Request{
			data: "test" + strconv.Itoa(i),
		})
	}
	fmt.Println(q.repo)
}
