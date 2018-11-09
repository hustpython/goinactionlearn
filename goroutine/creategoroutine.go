//这个示例程序展示如何创建goroutine
//以及调度行为

package main 
import(
	"runtime"
	"sync"
	"fmt"
)
func main(){
	//分配给一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)
	//wg用来等待程序完成
	//计数+2,表示要等待两个goruntine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")
	//申明一个匿名函数
	//并创建一个goroutine
	go func(){
		//在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		for count:=0;count<3;count++{
			for char:='a';char < 'a'+26;char++{
				fmt.Printf("%c ",char)
			}
		}
	}()
    //申明一个匿名函数
	//并创建一个goroutine
	go func(){
		//在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()
		for count:=0;count<3;count++{
			for char:='A';char < 'A'+26;char++{
				fmt.Printf("%c ",char)
			}
		}
	}()
	
	//等待goruntine介绍
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")


}