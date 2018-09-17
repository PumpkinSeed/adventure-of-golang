package loadbalancer

type Request struct {
	data interface{}
	resp chan interface{}
}

func NewRequest(data interface{}) Request {
	resp := make(chan interface{}, 1)
	return Request{
		data: data,
		resp: resp,
	}
}
