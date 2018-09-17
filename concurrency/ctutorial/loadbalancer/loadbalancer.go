package loadbalancer

const maxRequestPerWorker = 10

type Balancer struct {
	workerFunc func(req Request) interface{}

	workers []Worker

	maxRequestPerWorker int
}

func NewBalancer(workerFunc func(req Request) interface{}) Balancer {
	//queueMutex := sync.RWMutex{}
	return Balancer{
		workerFunc: workerFunc,
	}
}

func (b *Balancer) Add(r *Request) {

}

func (b *Balancer) SetMaxRequestPerWorker(n int) {
	b.maxRequestPerWorker = n
}

func (b *Balancer) Balance() {
	for {

	}
}
