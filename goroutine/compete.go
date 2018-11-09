//这个示例程序展示如何在程序里造成竞争状态
//实际上不希望出现这种情况
package main 

import (
	"fmt"
	"sync"
	"runtime"
)
var (
	//counter是所有goroutine都要增加其值的变量
	counter int
	wg sync.WaitGroup
)
func incCounter(id int){
	defer wg.Done()
	for count :=0;count <2 ;count++{
		//捕获counter的值
		value := counter 
		//当前goruntine从线程退出,并放回到队列
		runtime.Gosched()
		//增加本地value变量的值
		value++
		//将该值保存回counter
		counter = value
	}
}
func main(){
	//计数+2表示,要等待两个goroutine
	wg.Add(2)
	//创建两个goroutine
	go incCounter(1)
	go incCounter(2)
	//等待goroutine结束
	wg.Wait()
	fmt.Println("Final Counter: ",counter)
}