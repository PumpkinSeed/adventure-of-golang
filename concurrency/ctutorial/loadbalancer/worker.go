package loadbalancer

type Worker struct {
	idx     int
	wok     chan Request
	f       func(req Request) interface{}
	close   chan bool
	pending int
}

func NewWorker(idx int, f func(req Request) interface{}) *Worker {
	return &Worker{
		idx:     idx,
		wok:     make(chan Request),
		f:       f,
		close:   make(chan bool),
		pending: 0,
	}
}

func (w *Worker) Work() {
	for {
		select {
		case req := <-w.wok:
			w.pending++
			req.resp <- w.f(req)
			w.pending--
		case <-w.close:
			close(w.wok)
			return
		}
	}
}

func (w *Worker) Request(req Request) { w.wok <- req }
func (w *Worker) Pending() int        { return w.pending }
func (w *Worker) Close()              { w.close <- true }
