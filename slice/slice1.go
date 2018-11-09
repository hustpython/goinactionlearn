//切片的使用
package main

import (
	"fmt"
)
func foo(slice []int) [] int{
	return slice
}
func main(){
	//make方法
	//创建一个字符串切片
	//其长度和容量都是5个元素
	slicestrmake := make([]string,3,5)
	//切片字面量
	slicestrliterals := []string{"red","blue","green"}

	sliceintliterals := []int{10,20,30}
    //切片赋值
	slicestrmake[2] = slicestrliterals[2]
	//创建nil切片,长度为0,容量为0

	// var slicenil []int
	// slicenilmake := make([]int,0)
	// slicenilliterals := []int{}

	//使用切片创建切片
	sliceold := []int{10,20,30,40,50}
	//创建一个新切片,长度为2,容量为4个元素
	//对底层数据组容量是k的切片slice[i:j]来说
	//长度:j - i
	//容量:k - i
	slicenew := sliceold[1:3]//长度2,容量4
	//两个切片共享底层数组,修改其中一个切片,
	//另一个切片也能感知到
	slicenew[1] = 35

	//使用append向切片增加元素
	sliceappendold := []int{10,20,30,40,50}
	sliceappendnew := sliceappendold[1:3]
	sliceappendnew = append(sliceappendnew,60)

	//增加第三个元素,通过第三个元素限制
	//slicesource.容量为1.防止了slicesource通过append
	//修改了source的数据
	source := []string{"Apple","Orange","Plum","Banana","Grape"}
	slicesource := source[2:3:3]
	slicesource = append(slicesource,"Kiwi")
	//将一个切片增加到另一个切片
	s1 := []int{1,2,3}
	s2 := []int{3,4,5}
	//将两个切片追加在一起,并显示结果
	fmt.Printf("%v\n",append(s1,s2...))

	//迭代切片
	for index,value := range source{
		fmt.Printf("Index:%d  Value:%s\n",index,value)
	}

	//组合切片
	slicezuhe := [][]int{{10},{100,200}}
	slicezuhe[0] = append(slicezuhe[0],20)
	fmt.Println(slicezuhe)
	//传递切片
	slicecd := make([]int,1e6)
	slicecd =  foo(slicecd)

	fmt.Println(slicestrliterals[0])
	fmt.Println(sliceintliterals[1])
	fmt.Println(slicestrmake[2])
	fmt.Println(sliceold)
	fmt.Println(slicenew)
	fmt.Println(sliceappendnew)
	fmt.Println(sliceappendold)
	fmt.Println(source)
	fmt.Println(slicesource)
}