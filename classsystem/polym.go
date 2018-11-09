//这是一个使用接口展示多态行为的示例程序
package main

import (
	"fmt"
)

type notifier interface{
	notify()
}
type user struct{
	name string
	email string
}
func (u *user) notify(){
	fmt.Printf("Sending user email to %s<%s>\n",
	u.name,
    u.email)
}
//定义程序管理员
type admin struct{
	name string
	email string
}
func (a *admin)notify(){
	fmt.Printf("Sending user email to %s<%s>\n",
		a.name,
		a.email)
}

func main(){
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)
	// 创建一个 admin 值并传给 sendNotification
	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}
func sendNotification(n notifier) {
	 n.notify()
	 }