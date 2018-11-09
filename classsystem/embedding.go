//这是一个示例程序展示如何将一个类型嵌入到另一个类型
//以及内部类型和外部类型之间的关系
package main

import (
	"fmt"
)
type user struct {
 name string
 email string
 }

 // notify 实现了一个可以通过 user 类型值的指针
 // 调用的方法
 func (u *user) notify(){
	fmt.Printf("Sending user email to %s<%s>\n",
	u.name,
    u.email)
}
 // admin 代表一个拥有权限的管理员用户
 type admin struct {
	user // 嵌入类型
	level string
 }
func main(){
	ad := admin{
		user:user{
			name:"john smith",
			email:"john@yahoo.com",
		},
		level:"super",
	}
	//我们可以直接访问内部类型的方法
	ad.user.notify()
	//内部类型的方法也被提升到外部类型
	//如果外部类型实现了接口的方法
	//内部类型的实现就不会被提升
	ad.notify()
}