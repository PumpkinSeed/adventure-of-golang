package loadbalancer

import (
	"sync"
)

type Queue struct {
	repo  []*Request
	front int
	back  int
	len   int
	mutex *sync.RWMutex
}

func NewQueue(mutex *sync.RWMutex) *Queue {
	q := new(Queue)
	q.repo = make([]*Request, 1)
	q.front, q.back, q.len = 0, 0, 0
	q.mutex = mutex
	return q
}

func (q Queue) Len() int {
	return q.len
}

func (q Queue) Front() *Request {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	return q.repo[q.front]
}

func (q Queue) Back() *Request {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	return q.repo[q.dec(q.back)]
}

func (q *Queue) PushFront(r *Request) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.lazyGrow()
	q.front = q.dec(q.front)
	q.repo[q.front] = r
	q.len++
}

func (q *Queue) PushBack(r *Request) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.lazyGrow()
	q.repo[q.back] = r
	q.back = q.inc(q.back)
	q.len++
}

func (q *Queue) PopFront() *Request {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.empty() {
		return nil
	}
	r := q.repo[q.front]
	q.repo[q.front] = nil
	q.front = q.inc(q.front)
	q.len--
	q.lazyShrink()
	return r
}

func (q *Queue) PopBack() *Request {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.empty() {
		return nil
	}
	q.back = q.dec(q.back)
	r := q.repo[q.back]
	q.repo[q.back] = nil
	q.len--
	q.lazyShrink()
	return r
}

func (q *Queue) resize(size int) {
	adjusted := make([]*Request, size)
	if q.front < q.back {
		copy(adjusted, q.repo[q.front:q.back])
	} else {
		n := copy(adjusted, q.repo[q.front:])
		copy(adjusted[n:], q.repo[:q.back])
	}
	q.repo = adjusted
	q.front = 0
	q.back = q.len
}

func (q *Queue) lazyGrow() {
	if q.full() {
		q.resize(len(q.repo) * 2)
	}
}

func (q *Queue) lazyShrink() {
	if q.sparse() {
		q.resize(len(q.repo) / 2)
	}
}

func (q Queue) full() bool    { return q.len == len(q.repo) }
func (q Queue) inc(i int) int { return (i + 1) & (len(q.repo) - 1) }
func (q Queue) dec(i int) int { return (i - 1) & (len(q.repo) - 1) }
func (q Queue) sparse() bool  { return 1 < q.len && q.len < len(q.repo)/4 }
func (q Queue) empty() bool   { return q.len == 0 }
