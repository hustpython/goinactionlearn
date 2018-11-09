package main 

import "fmt"

const a = iota
const b = iota 

func main(){
	fmt.Println("a的常量值为:")
	fmt.Println(a)

	fmt.Println("b的常量值为:")
	fmt.Println(b)
}