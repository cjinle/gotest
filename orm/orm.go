package mylib

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	Id       int
	Username string
	Age      int
}

func Orm() {

	orm.RegisterModel(new(UserInfo))

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8")

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	// user := new(User)
	// user.Age = 10
	// user.Username = "slene111"

	// // fmt.Println(o.Insert(user))
	// var err error
	// var id int64
	// id, err = o.Insert(user)

	// user2 := &User{Id: int(id)}
	// err = o.Read(user2)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user2)

	user := new(UserInfo)
	user.Id = 5
	user.Age = 18
	user.Username = "chenjinle"

	num, err := o.Update(user)
	fmt.Println(num)
	fmt.Println(err)
}
