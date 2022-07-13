package ref

import (
	"fmt"
	"reflect"
)

func Test() {
	x := 3.14
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value", reflect.ValueOf(x))

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	//Value类型提供了Int(),Float()等方法，可以让我们获取存在里面的值
	fmt.Println("value", v.Float())

	t := reflect.TypeOf(x)
	fmt.Println("v", reflect.TypeOf(v))
	fmt.Println("t", reflect.TypeOf(t))

	//从反射对象到接口变量
	i := v.Interface()
	fmt.Println("i:", reflect.TypeOf(i), reflect.ValueOf(i))
}
