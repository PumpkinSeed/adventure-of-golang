package loadbalancer

import (
	"log"
	"sync"
	"time"
)

const maxRequestPerWorker = 10

// @TODO add context with cancel

type Balancer struct {
	workerFunc          func(req Request) interface{}
	workers             []*Worker
	requestChan         chan Request
	maxRequestPerWorker int
	idState             int
	workersMutex        *sync.RWMutex
	debug               bool
}

func NewBalancer(workerFunc func(req Request) interface{}) Balancer {
	return Balancer{
		workerFunc:          workerFunc,
		idState:             0,
		workersMutex:        &sync.RWMutex{},
		maxRequestPerWorker: maxRequestPerWorker,
		requestChan:         make(chan Request),
	}
}

func (b *Balancer) Add(data interface{}) chan interface{} {
	var req = NewRequest(data)
	b.requestChan <- req
	return req.resp
}

func (b *Balancer) SetMaxRequestPerWorker(n int) {
	b.maxRequestPerWorker = n
}

func (b *Balancer) Balance() {
	b.scale()
	timeTick := time.Tick(10 * time.Millisecond)
	for {
		select {
		case req := <-b.requestChan:
			b.DebugLog("handle request")
			b.handleRequest(req)
		case <-timeTick:
			b.scale()
			b.clean()
		}
	}
}

func (b *Balancer) Interrupt() {
	close(b.requestChan)
	for _, worker := range b.workers {
		if worker.IsRunning() {
			worker.Close()
		}
	}
}

func (b *Balancer) SetDebug(debug bool) {
	b.debug = debug
}

func (b *Balancer) NumOfWorkers() int {
	return len(b.workers)
}

func (b Balancer) DebugLog(msg string) {
	if b.debug {
		log.Println(msg)
	}
}

func (b *Balancer) handleRequest(req Request) {
	for _, worker := range b.workers {
		if worker.IsRunning() && worker.Pending() < b.maxRequestPerWorker {
			worker.Request(req)
		} else {
			b.scale()
			time.Sleep(100 * time.Millisecond)
			// @TODO fix this
			b.handleRequest(req)
		}
	}
}

func (b *Balancer) clean() {
	for i, worker := range b.workers {
		if !worker.IsRunning() && worker.Pending() < 1 {
			worker.Close()
			b.workersMutex.Lock()
			b.workers = append(b.workers[:i], b.workers[i+1:]...)
			b.workersMutex.Unlock()
		}
	}
}

func (b *Balancer) scale() {
	if len(b.workers) < 1 {
		b.addWorker()
	}
	var available bool
	for _, worker := range b.workers {
		if worker.IsRunning() && worker.Pending() < b.maxRequestPerWorker {
			available = true
		} else if !worker.IsRunning() {
			worker.SetPending(0)
		}
	}
	if !available {
		b.addWorker()
	}
}

func (b *Balancer) addWorker() {
	b.workersMutex.Lock()
	defer b.workersMutex.Unlock()
	w := NewWorker(b.idState, b.workerFunc)
	b.idState++
	go w.Work()
	b.workers = append(b.workers, w)
}
