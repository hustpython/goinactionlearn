package main 

import (
	"fmt"
)

//notifier 是一个定义了
//通知类行为的接口
type notifier interface{
	notify()
}
//user在程序里定义一个用户类型
type user struct{
	name string
	email string
}
//notify是使用指针接收者实现的方法
func(u* user) notify(){
	fmt.Printf("Sending user email to %s<%s>\n",
	u.name,
u.email)
}
func main(){
	u:=user{"Bill","bill@email.com"}
	//如果使用指针接收者来实现一个接口，那么只有指向那个
	//类型的指针才能够实现对应的接口。如果使用值
	//接收者来实现一个接口，那么那个类型的值和指针
	//都能够实现对应的接口
	//为什么这里是&u的原因
	sendNotification(&u)
}
func sendNotification(u notifier){
	u.notify()
}