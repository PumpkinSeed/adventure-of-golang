package loadbalancer

type Request struct {
	data interface{}
	resp chan interface{}
}

func NewRequest(b Balancer, data interface{}) chan interface{} {
	resp := make(chan interface{}, 1)

	return resp
}
