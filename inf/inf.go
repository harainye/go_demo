package inf

import (
	"encoding/json"
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

type Iphone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (iphone Iphone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p Person) test() {
	fmt.Println("testing Person,name=", p.Name)
}

func (p Person) calculate(n int) int {
	res := 0
	for i := 0; i < n; i++ {
		res += i
	}
	return res
}

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}
func (i *integer) change() {
	*i += 1
}

func CallPhone() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(Iphone)
	phone.call()
	person := Person{"tom", 30}

	jsonStr, err := json.Marshal(person)
	if err != nil {
		fmt.Println("json处理错误", err)
	}

	fmt.Println("jsonStr:", string(jsonStr))

	var p Person
	p.Name = "Micheal"
	p.test()

	n := 10
	res := p.calculate(n)
	fmt.Println("res:", res)

	var i integer = 10
	i.print()
	i.change()
	fmt.Println("i=", i)

}
