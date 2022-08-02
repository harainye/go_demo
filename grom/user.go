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

	/*user := User{UserName: "test", Age: 25, Addr: "福州", Pic: "/static/img.png", DisplayName: "test"}

	// 插入(创建)数据
	err = conn.Create(&user).Error
	if err != nil {
		fmt.Println("insert err:", err)
		return
	}*/
	user := &User{Id: 1}

	// 更新单个字段
	// conn.Model(user).Update("UserName", "lisi")
	// 更新多个字段
	conn.Model(user).Update(map[string]interface{}{"UserName": "lisi", "DisplayName": "李四"})

	// 查询单个
	user = new(User)
	conn.First(user, 1)
	fmt.Println(user)

	// 删除
	conn.Delete(user)

	// 查询列表
	var users []User
	conn.Find(&users)
	fmt.Println(users)

}
