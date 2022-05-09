package simpledi

import (
	"fmt"
	"reflect"
)

func Put[T any](c IContainer, instance T) {
	PutWithName[T](c, instance, "")
}

func PutWithName[T any](c IContainer, instance T, instanceName string) {
	interfaceType := reflect.TypeOf((*T)(nil)).Elem()
	c.Put(getInterfaceKey(interfaceType), instanceName, reflect.ValueOf(instance))
}

func Get[T any](c IContainer) T {
	return GetWithName[T](c, "")
}

func GetWithName[T any](c IContainer, instanceName string) T {
	key := getInterfaceKey(reflect.TypeOf((*T)(nil)).Elem())
	return c.Get(key, instanceName).Interface().(T)
}

func Inject(c IContainer, instancePtrToInject any) {
	targetType := reflect.TypeOf(instancePtrToInject)
	if targetType.Kind() != reflect.Ptr {
		panic(fmt.Errorf("simpledi: inject target <%s> is not ptr", targetType.String()))
	}
	targetElemType := targetType.Elem()
	targetElemValue := reflect.ValueOf(instancePtrToInject).Elem()
	targetFieldCount := targetElemType.NumField()
	for i := 0; i < targetFieldCount; i++ {
		fieldType := targetElemType.Field(i)
		instanceName, ok := fieldType.Tag.Lookup("inject")
		if !ok {
			continue
		}
		interfaceKey := getInterfaceKey(fieldType.Type)
		targetElemValue.FieldByIndex(fieldType.Index).Set(c.Get(interfaceKey, instanceName))
	}
}

// ---

func getInterfaceKey(interfaceType reflect.Type) string {
	return fmt.Sprintf("%s.%s", interfaceType.PkgPath(), interfaceType.Name())
}
