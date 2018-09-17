package loadbalancer

import (
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	workerFunc := func(req Request) interface{} {
		time.Sleep(100 * time.Millisecond)
		return req.data
	}
	w := NewWorker(1234, workerFunc)

	go w.Work()

	time.Sleep(100 * time.Millisecond)

	respChan := make(chan interface{}, 1)
	req := Request{
		data: "1234",
		resp: respChan,
	}

	w.Request(req)
	time.Sleep(10 * time.Millisecond)
	pending := w.Pending()
	if pending != 1 {
		t.Errorf("Pending should be 1, instead of %d", pending)
	}
	resp := <-respChan
	if resp.(string) != "1234" {
		t.Errorf("Worker should return '1234', instead of %s", resp.(string))
	}

	w.Close()
}
