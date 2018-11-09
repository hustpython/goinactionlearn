//这个示例程序展示如何使用atomic包来提供
//对数值的安全访问

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}
func incCounter(id int) {
	//在函数退出时调用Done来通知main函数工作已完成
	defer wg.Done()
	for count := 0; count < 2; count++ {
		//安全地对counter+1
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
