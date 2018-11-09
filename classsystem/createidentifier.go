//这是一个示例程序展示无法从另一个包
//里访问为公开的标识符(开头小写)
package main 

import(
	"fmt"
	"github.com/goinaction/code/chapter5/listing68/counters"
)
func main(){
	//counter := counters.alertCounter(10)
	counter := counters.New(10)
	fmt.Printf("Counter: %d\n",counter)
}