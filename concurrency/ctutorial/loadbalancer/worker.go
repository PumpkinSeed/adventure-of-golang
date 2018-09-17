package loadbalancer

type Worker struct {
	idx     int
	wok     chan Request
	f       func(req Request) interface{}
	close   chan bool
	pending int
	running bool
}

func NewWorker(idx int, f func(req Request) interface{}) *Worker {
	return &Worker{
		idx:     idx,
		wok:     make(chan Request),
		f:       f,
		close:   make(chan bool),
		pending: 0,
		running: false,
	}
}

func (w *Worker) Work() {
	w.running = true
	for {
		select {
		case req := <-w.wok:
			w.pending++
			req.resp <- w.f(req)
			w.pending--
		case <-w.close:
			close(w.wok)
			w.running = false
			return
		}
	}
}

func (w *Worker) Request(req Request) { w.wok <- req }
func (w Worker) Pending() int         { return w.pending }
func (w *Worker) Close()              { w.close <- true }
func (w Worker) IsRunning() bool      { return w.running }
func (w *Worker) SetPending(n int)    { w.pending = n }
