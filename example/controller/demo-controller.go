package controller

import (
	"fmt"
	"simpledi/example/service"
)

type DemoController struct {
	HelloSvc         service.IHelloService   `inject:""`
	MessageEmptySvc  service.IMessageService `inject:""`
	MessageBananaSvc service.IMessageService `inject:"Banana"`
}

func (d DemoController) Visit() {
	fmt.Printf("Hello: [%s] \n", d.HelloSvc.SayHello())
	fmt.Printf("Message from Empty: [%s] \n", d.MessageEmptySvc.Message())
	fmt.Printf("Message from Banana: [%s] \n", d.MessageBananaSvc.Message())
}
