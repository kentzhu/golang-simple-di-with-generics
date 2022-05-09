package main

import (
	"simpledi"
	"simpledi/example/controller"
	"simpledi/example/service"
	"simpledi/example/service/impl"
)

func main() {

	// create a di instance container
	container := simpledi.NewContainer()

	// put instances into the container
	simpledi.Put[service.IHelloService](container, impl.NewHelloService())
	simpledi.Put[service.IMessageService](container, impl.NewMessageService("My name is empty"))
	simpledi.PutWithName[service.IMessageService](container, impl.NewMessageService("My name is Banana"), "Banana")

	// demo target instance to inject
	// for example: target could be the Controller in mvc web application
	ctl := &controller.DemoController{}
	simpledi.Inject(container, ctl)
	ctl.Visit()
}
