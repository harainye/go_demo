package arr

import "fmt"

//值传递只能修改形参这个局部变量的值，而无法修改实参的值
func change(array [5]int) {
	fmt.Println(&array[0])
	for index, value := range array {
		array[index] = -value
	}
	fmt.Println(array)
}

//传入指针才能真正修改实参的值
func changeP(array *[5]int) {
	fmt.Println((*array)[1])
	for index, value := range *array {
		(*array)[index] = -value
	}
}

func changeD(array []int) {
	for index, value := range array {
		array[index] = -value
	}
}

// 数组值传递修改
func DoArr() {
	//固定长度数组
	array1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array1)
	for index, value := range array1 {
		fmt.Println(index, value)
	}
	fmt.Println(&array1[0])
	change(array1)
	fmt.Println(array1)
	changeP(&array1)
	fmt.Println(array1)
}

func DiArr() {
	//声明但没分配空间,类似于c语言中的 int *p
	var array1 []int
	fmt.Println(array1)
	//声明并分配3个单元的空间
	var array2 = []int{12, 33, 45}
	fmt.Println(array2)
	//使用:=声明
	array3 := []int{33, 44, 55, 66}
	fmt.Println(array3)
	//使用make函数开辟空间,类似于c语言中的malloc函数，或者c++中的auto p = new int[5]
	array4 := make([]int, 5)
	fmt.Println(array4)

	//动态数组
	var array5 = []int{12, 33, 45}
	fmt.Println(array5)
	changeD(array5)
	fmt.Println(array5)
}
