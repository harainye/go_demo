package grom

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //"_" 代码不直接使用包, 底层链接要使用!
	"github.com/jinzhu/gorm"
)

type User struct {
	Id          int
	UserName    string
	DisplayName string
	Age         int
	Addr        string
	Pic         string
}

type Job struct {
	Id          int
	Name        string
	UserId      int
	Description string
}

type Role struct {
	Id          int
	Name        string
	DisplayName string
	Age         int
	Addr        string
	Pic         string
}

func DoGrom() {

	// mysql: 数据库的驱动名
	// 链接数据库 --格式: 用户名:密码@协议(IP:port)/数据库名？xxx&yyy&
	conn, err := gorm.Open("mysql", "root:123abc,./@tcp(192.168.9.75:3306)/go")
	if err != nil {
		fmt.Println("gorm.Open err:", err)
		return
	}

	// gorm查找struct名对应数据库中的表名的时候会默认把你的struct中的大写字母转换为小写并加上“s”
	// db.SingularTable(true) 让grom转义struct名字的时候不用加上s
	conn.SingularTable(true)

	defer conn.Close()
	// 先创建数据 --- 创建对象

	/*user := User{UserName: "test1", Age: 30, Addr: "厦门", Pic: "/static/img.png", DisplayName: "test"}

	// 插入(创建)数据
	err = conn.Create(&user).Error
	if err != nil {
		fmt.Println("insert err:", err)
		return
	}*/
	// 根据Map创建
	/*err = conn.Model(&User{}).Create(map[string]interface{}{
		"UserName":    "张三",
		"Age":         20,
		"Addr":        "厦门",
		"Pic":         "/static/img.png",
		"DisplayName": "张三",
	}).Error*/

	user := &User{Id: 2}

	// 更新单个字段
	// conn.Model(user).Update("UserName", "lisi")
	// 更新多个字段
	conn.Model(user).Update(map[string]interface{}{"UserName": "lisi", "DisplayName": "李四"})

	// 查询单个
	user = new(User)
	conn.First(&user)
	fmt.Println("====查询第一个用户", user)

	result := conn.Take(&user)
	fmt.Println("====获取一条记录，没有指定排序字段", user)
	fmt.Println("====获取的记录数：", result.RowsAffected)
	// 删除
	//conn.Delete(user)

	// 查询列表
	var users []User
	//conn.Find(&users)
	//fmt.Println(users)

	//conn.Find(&users, []int{1, 2, 3})
	//fmt.Println("====用in查询", users)

	err = conn.Model(&User{}).Select("user.user_name, job.name, job.description").Joins("left join job on job.user_id = user.id").Scan(&Job{}).Error
	// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id
	/*rows, err := conn.Table("user").Select("user.id, user.user_name, user.age").Joins("left join job on job.user_id = user.id").Rows()
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.UserName, &u.Age)
		if err != nil {
			panic(err)
		}
		users = append(users, u)
	}*/
	fmt.Println("====关联查询", users)

	var job Job
	conn.Raw("SELECT * FROM job where user_id =2").Scan(&job)
	fmt.Println("====语句查询", job)
}
