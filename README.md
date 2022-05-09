# Golang Simple DI with Generics 基于泛型的简单DI实现

You can see a simple example at `\example`

你可以在`\example`看到简单示例:


```go
// create a di instance container
// 创建一个DI容器
container := simpledi.NewContainer()

// put instances into the container
// 将被依赖的实例放入DI
simpledi.Put[service.IHelloService](container, impl.NewHelloService())
simpledi.Put[service.IMessageService](container, impl.NewMessageService("My name is empty"))
simpledi.PutWithName[service.IMessageService](container, impl.NewMessageService("My name is Banana"), "Banana")

// demo target instance to inject
// for example: target could be the Controller in mvc web application
// 使用依赖注入的简单演示
// 注入目标可以是MVC结构Web应用中的控制器
ctl := &controller.DemoController{}
simpledi.Inject(container, ctl)
ctl.Visit()
```

The target instance is simple too. Just add tag `inject`.
注入的容器也非常简单。只要加上Tag `inject`即可。

```go
type DemoController struct {
	HelloSvc         service.IHelloService   `inject:""`
	MessageEmptySvc  service.IMessageService `inject:""`
	MessageBananaSvc service.IMessageService `inject:"Banana"`
}
```


---

# reference 参考

https://github.com/zekroTJA/di
