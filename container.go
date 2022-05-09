package simpledi

import (
	"fmt"
	"reflect"
)

type IContainer interface {
	Get(key, instanceName string) reflect.Value
	Put(key, instanceName string, instance reflect.Value)
}

func NewContainer() IContainer {
	return &container{
		m: map[string]map[string]reflect.Value{},
	}
}

type container struct {
	m map[string]map[string]reflect.Value
}

func (c *container) Get(key, instanceName string) reflect.Value {
	if _, ok := c.m[key]; !ok {
		panic(fmt.Errorf("simpledi: instance [%s] has not registered", key))
	}
	if _, ok := c.m[key][instanceName]; !ok {
		panic(fmt.Errorf("simpledi: instance [%s] name [%s] has not registered", key, instanceName))
	}
	return c.m[key][instanceName]
}

func (c *container) Put(key, instanceName string, instance reflect.Value) {
	if _, ok := c.m[key]; !ok {
		c.m[key] = map[string]reflect.Value{}
	}
	c.m[key][instanceName] = instance
}
