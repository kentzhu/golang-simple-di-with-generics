package impl

import "simpledi/example/service"

func NewHelloService() service.IHelloService {
	return &helloService{}
}

type helloService struct {
}

func (m *helloService) SayHello() string {
	return "Hello"
}
