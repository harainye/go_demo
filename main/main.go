package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_projects/gintest"
	"go_projects/redis"
	"log"
	"net/http"
	"net/mail"
	"reflect"
	"regexp"
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

	/*m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(router, &m))*/

	//router.Run(":8888")

	//fmt.Println(regexp.IsEmailValid("test44@qq.com"))
	//fmt.Println(valid("test44@qq.com"))
	//fmt.Println(CheckMobile("15059459039"))
	fmt.Println(CheckDate("2022-09-30"))
	//fmt.Println(CheckIdCard("35081246682578451X"))
	//fmt.Println(CheckPhone("15059459039"))
	//fmt.Println(CheckEmail("aa_ss22@163.com"))
	//fmt.Println(checkUri("user/88/964488/2022/1008/203557p6qyXsau.jpeg"))
	fmt.Println(checkSpecial("+86 0597111"))
	//fmt.Println("demo return:", Demo())
	//fmt.Println("demo2 return:", Demo2())
	//rec.RunRecover()
}
func Demo() int {
	// 实际上return 执行了两步操作。
	//因为返回值没有命名，所以return 之前
	//首先默认创建了一个临时零值变量(假设为s)作为返回值
	//然后将i赋值给s,此时s的值是0。后续的操作是针对i进行的，
	//所以不会影响s, 此后因为s不会更新，
	//所以return s 不会改变
	//    相当于：
	//          var i int
	//          s := i
	//          return s
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}
func Demo2() (i int) {
	//因为返回值已经提前定义了，不会产生临时零值变量，
	//返回值就是提前定义的变量，后续所有的操作也都是基于已经定义的变量，
	//任何对于返回值变量的修改都会影响到返回值本身。
	//
	//就相当于s就是命名的变量i, 后续所有的操作都是基于
	//命名变量i(s),返回值也是i, 所以每一次defer操作，
	//都会更新返回值i。
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i // 或者直接 return 效果相同
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// CheckMobile 检验手机号
func CheckEmail(email string) bool {
	// 匹配规则
	// $ 结束符
	regRuler := "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	// 正则调用规则
	reg := regexp.MustCompile(regRuler)
	// 返回 MatchString 是否匹配
	return reg.MatchString(email)

}

// CheckMobile 检验手机号
func CheckMobile(phone string) bool {
	// 匹配规则
	// ^1第一位为一
	// [345789]{1} 后接一位345789 的数字
	// \\d \d的转义 表示数字 {9} 接9位
	// $ 结束符
	regRuler := "^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\\d{8}$"
	// 正则调用规则
	reg := regexp.MustCompile(regRuler)
	// 返回 MatchString 是否匹配
	return reg.MatchString(phone)

}

func CheckDate(phone string) bool {
	regRuler := "((((19|20)\\d{2})-(0?(1|[3-9])|1[012])-(0?[1-9]|[12]\\d|30))|(((19|20)\\d{2})-(0?[13578]|1[02])-31)|(((19|20)\\d{2})-0?2-(0?[1-9]|1\\d|2[0-8]))|((((19|20)([13579][26]|[2468][048]|0[48]))|(2000))-0?2-29))$"
	// 正则调用规则
	reg := regexp.MustCompile(regRuler)
	// 返回 MatchString 是否匹配
	return reg.MatchString(phone)
}

func checkUri(uri string) bool {
	regRuler := "^[A-Za-z0-9_\\./]+$"
	// 正则调用规则
	reg := regexp.MustCompile(regRuler)
	// 返回 MatchString 是否匹配
	return reg.MatchString(uri)
}

// CheckIdCard 检验身份证
func CheckIdCard(card string) bool {
	//18位身份证 ^(\d{17})([0-9]|X)$
	// 匹配规则
	// (^\d{15}$) 15位身份证
	// (^\d{18}$) 18位身份证
	// (^\d{17}(\d|X|x)$) 18位身份证 最后一位为X的用户
	regRuler := "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)
	// 返回 MatchString 是否匹配
	return reg.MatchString(card)
}

// CheckPhone 检验电话号码
func CheckPhone(phone string) bool {
	// 国内电话号码(0511-4405222、021-87888822)
	// 匹配规则
	// \d{3}-\d{8}|\d{4}-\d{7}
	//regRuler := "^((0\\d{2,3})-)(\\d{7,8})(-(\\d{3,}))?$"
	regRuler := "/\\d{3}-\\d{8}|\\d{4}-\\d{7}|^(13[0-9]|14[01456879]|15[0-3,5-9]|16[2567]|17[0-8]|18[0-9]|19[0-3,5-9])\\d{8}$/"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)
	// 返回 MatchString 是否匹配
	return reg.MatchString(phone)
}

func checkSpecial(specialStr string) bool {
	regRuler := "^[0-9_@#\\- $%!~*()&,;=+]+$"
	// 正则调用规则
	reg := regexp.MustCompile(regRuler)
	// 返回 MatchString 是否匹配
	return reg.MatchString(specialStr)
}

//binding type interface 要修改的结构体
//value type interace 有数据的结构体
func structAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		fmt.Printf("====type:", vTypeOfT.Field(i).Type)
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
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
