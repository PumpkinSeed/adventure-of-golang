package loadbalancer

type Request struct {
	data interface{}
	resp chan interface{}
}

func NewRequest(data interface{}) chan interface{} {

	return nil
}
