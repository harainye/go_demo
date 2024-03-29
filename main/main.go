package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_defer "go_projects/defer"
	"go_projects/gintest"
	"go_projects/redis"
	"go_projects/routine"
	"log"
	"net/http"
	"reflect"
)

// 结构体
type ss struct {
	int
	string
	bool
	float64
}

func (s ss) Method1(i int) string  { return "结构体方法1" }
func (s *ss) Method2(i int) string { return "结构体方法2" }

var (
	structValue = ss{ // 结构体
		20,
		"结构体",
		false,
		64.0,
	}
)

// 复杂类型
var complexTypes = []interface{}{
	structValue, &structValue, // 结构体
	structValue.Method1, structValue.Method2, // 方法
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

type Car struct {
	Name  string
	Price float32
}

type Person struct {
	Name string
	Age  int
}

type Teacher struct {
	Person
	ProjectName string
}

type Student struct {
	Name  string
	Age   int
	Grade int
}

type User struct {
	Name  string
	Age   int
	Sex   int
	Grade int
}

type ClassService interface {
	Class()
}

func (t *Teacher) Class() {
	fmt.Printf("老师:%s，教的科目是: %s \n", t.Name, t.ProjectName)
}

func (s *Student) Class() {
	fmt.Printf("学生：%s，现在的年级是：%d\n", s.Name, s.Grade)
}

func Shangke(class ClassService) {
	class.Class()
}

type ElectricCar struct {
	Car
	Battery int32
}
type PetrolCar struct {
	Car
	Gasoline int32
}

//定义一个接口
type RunService interface {
	Run()
}

// 实现1
func (car *PetrolCar) Run() {
	fmt.Printf("%s PetrolCar run \n", car.Name)
}

// 实现2
func (car *ElectricCar) Run() {
	fmt.Printf("%s ElectricCar run \n", car.Name)
}

func Do(run RunService) {
	run.Run()
}

func main() {
	/*src := []byte{1, 2, 3, 4}
	var dest []byte = make([]byte, len(src))
	copy(dest, src)
	dest[0] = 8
	fmt.Println(dest)
	fmt.Println(src)*/
	/*xp := ElectricCar{
		Car{Name: "xp", Price: 200},
		70,
	}
	petrolCar := PetrolCar{
		Car{Name: "BMW", Price: 300},
		50,
	}
	Do(&xp)
	Do(&petrolCar)*/

	//teacher := Teacher{Person{Name: "王老师", Age: 30}, "English"}
	//
	//student := Student{Person{Name: "小明", Age: 10}, 3}

	//teacher.Class()
	//student.Class()
	//Shangke(&teacher)
	//Shangke(&student)
	//arr.DiArr()
	//grom.DoGrom()
	// 测试复杂类型
	/*for i := 0; i < len(complexTypes); i++ {
		PrintInfo(complexTypes[i])
	}*/
	/*filePath := "D:\\JIRA/jira.sql" //找一个大的文件，如日志文件
	start := time.Now()
	file.Read1(filePath)
	t1 := time.Now()
	fmt.Printf("Cost time %v\n", t1.Sub(start))
	file.Read2(filePath)
	t2 := time.Now()
	fmt.Printf("Cost time %v\n", t2.Sub(t1))
	file.Read3(filePath)
	t3 := time.Now()
	fmt.Printf("Cost time %v\n", t3.Sub(t2))*/

	//tcp.TcpCo()

	//template.RenderTemplate()
	//template.RenderHTemplate()
	//ref.Test()
	//mongodb.ConnMongo()
	//rate.DoRate()

	//inf.CallPhone()

	//vip.GetIniFile()
	//vip.GetJSONFile()
	//vip.GetYamlFile()

	// defer demo
	//_defer.DeferCall()
	//_defer.DeferCall2()
	_defer.DeferCall3()

	// goroutine
	//routine.TestChan()
	routine.TestSelect()

	// test lock
	//locktest.TestLock()

	redis.DoRedis()

	log.Println("开始启动服务...")
	router := gin.Default()

	// 强制日志颜色化
	gin.ForceConsoleColor()

	//router.Use(gintest.Logger())

	// gin test
	ginV1 := router.Group("/v1/gintest")
	{
		ginV1.POST("/post", gintest.Post)
		ginV1.POST("/form_post", gintest.FormPost)
		ginV1.GET("/json", gintest.JSON)
		ginV1.GET("/pureJson", gintest.PureJSON)
		ginV1.POST("/upload", gintest.Upload)
		ginV1.GET("/hasgo", gintest.HasGo)
		ginV1.GET("/nogo", gintest.NoGo)
		ginV1.GET("/uri/:name/:id", gintest.Uri)
		ginV1.GET("/redirect", func(c *gin.Context) {
			c.Request.URL.Path = "/v1/gintest/json"
			router.HandleContext(c)
		})
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // c.Request.URL.Query().Get("lastname") 的一种快捷方式
		example := c.MustGet("example").(string)

		c.String(http.StatusOK, "Hello %s %s %s", firstname, lastname, example)
	})

	router.Run(":8888")

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func PrintInfo(i interface{}) {
	if i == nil {
		fmt.Println("--------------------")
		fmt.Printf("无效接口值：%v\n", i)
		fmt.Println("--------------------")
		return
	}
	v := reflect.ValueOf(i)
	PrintValue(v)
}

func PrintValue(v reflect.Value) {
	fmt.Println("--------------------")
	// ----- 通用方法 -----
	fmt.Println("String             :", v.String())  // 反射值的字符串形式
	fmt.Println("Type               :", v.Type())    // 反射值的类型
	fmt.Println("Kind               :", v.Kind())    // 反射值的类别
	fmt.Println("CanAddr            :", v.CanAddr()) // 是否可以获取地址
	fmt.Println("CanSet             :", v.CanSet())  // 是否可以修改
	if v.CanAddr() {
		fmt.Println("Addr               :", v.Addr())       // 获取地址
		fmt.Println("UnsafeAddr         :", v.UnsafeAddr()) // 获取自由地址
	}
	// 获取方法数量
	fmt.Println("NumMethod          :", v.NumMethod())
	if v.NumMethod() > 0 {
		// 遍历方法
		i := 0
		for ; i < v.NumMethod()-1; i++ {
			fmt.Printf("    ┣ %v\n", v.Method(i).String())
			//            if i >= 4 { // 只列举 5 个
			//                fmt.Println("    ┗ ...")
			//                break
			//            }
		}
		fmt.Printf("    ┗ %v\n", v.Method(i).String())
		// 通过名称获取方法
		fmt.Println("MethodByName       :", v.MethodByName("String").String())
	}

	switch v.Kind() {
	// 结构体：
	case reflect.Struct:
		fmt.Println("=== 结构体 ===")
		// 获取字段个数
		fmt.Println("NumField           :", v.NumField())
		if v.NumField() > 0 {
			var i int
			// 遍历结构体字段
			for i = 0; i < v.NumField()-1; i++ {
				field := v.Field(i) // 获取结构体字段
				fmt.Printf("    ├ %-8v %v\n", field.Type(), field.String())
			}
			field := v.Field(i) // 获取结构体字段
			fmt.Printf("    └ %-8v %v\n", field.Type(), field.String())
			// 通过名称查找字段
			if v := v.FieldByName("ptr"); v.IsValid() {
				fmt.Println("FieldByName(ptr)   :", v.Type().Name())
			}
			// 通过函数查找字段
			v := v.FieldByNameFunc(func(s string) bool { return len(s) > 3 })
			if v.IsValid() {
				fmt.Println("FieldByNameFunc    :", v.Type().Name())
			}
		}
	}
}
