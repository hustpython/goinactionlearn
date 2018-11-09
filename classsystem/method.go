package main

import (
	"fmt"
)

type user struct{
	name string
	email string
}
//notify使用值接收者实现一个方法
//func和函数名之间的参数被称作接收者,如果一个函数有接收者
//这个函数就被称为方法
//值接收者使用值的副本来调用方法
func (u user) notify(){
	fmt.Printf("Sending User Email To %s<%s>\n",
	u.name,
    u.email)
}  
//指针接收者使用实际的值来调用方法,
//调用对值的修改会,立即反映出来
func (u *user)changeEmail(email string){
	u.email = email
} 
func main(){
	bill:=user{"Bill","bill@email.com"}
	//值接收者
	bill.notify()
    //使用user类型的指针也可以用来调用
	lisa := &user{"Lisa","lisa@email.com"}
	//指针接收者
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}