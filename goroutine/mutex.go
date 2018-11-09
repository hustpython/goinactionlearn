// 这个示例程序展示如何使用互斥锁
// 来定义一段需要同步访问的代码临界区
// 资源的同步访问

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	// mutex 用来定义一段代码临界区
	mutex sync.Mutex
)

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}
func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goruntine进入这个临界区
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
		// 释放锁,允许其他正在等待的goruntine进入临界区
	}
}
