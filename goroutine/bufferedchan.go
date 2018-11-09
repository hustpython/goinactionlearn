// 这个示例程序展示如何使用
// 有缓冲的通道和固定数目的
// goruntine来处理一堆工作
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 要使用的goroutine的数量
	taskLoad         = 10 // 要处理的工作的数量
)

var wg sync.WaitGroup

// init 初始化包,Go 语言运行时会在其他代码执行之前
// 优先执行这个函数
func init() {
	// 初始化随机种子
	rand.Seed(time.Now().Unix())
}

func main() {
	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)
	// 启动goruntine来处理工作
	wg.Add(numberGoroutines)
	//  Go 语言的调度器带有随机成分
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		// 可以用`Sprintf`来将格式化后的字符串赋值给一个变量
		tasks <- fmt.Sprintf("Task: %d", post) // 通过改变tasks触发worker函数
	}
	// 当所有的工作都处理完时关闭通道
	// 以便所有的goroutine退出
	close(tasks)
	wg.Wait()
}
func worker(tasks chan string, worker int) {
	// 通知函数已经返回
	defer wg.Done()
	for {
		// 等待分配工作
		task, ok := <-tasks
		if !ok {
			// 这意味着通道已经空了,并且已被关闭
			fmt.Printf("Worker:%d: Shuting Down\n", worker)
			return
		}
		// 显示我们开始工作了
		fmt.Printf("Worker: %d : Started %s\n", worker, task)
		// 随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		//显示我们完成了工作
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
