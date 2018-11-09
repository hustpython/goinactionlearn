// 这个示例程序展示来自不同标准库的不同函数
// 是如何使用io.Writer接口的

package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 创建一个Buffer值,并将一个字符串写入Buffer
	// 使用io.Writer的write方法
	var b bytes.Buffer
	b.Write([]byte("Hello"))

	// 使用Fprintf来将一个字符串拼接到Buffer里
	// 将bytes.Buffer的地址作为io.Writer类型的传入
	fmt.Fprintf(&b, "World!\n")
	// 将Buffer的内容输出到标准输出设备
	// 将os.File值的地址作为io.Write类型的传入
	b.WriteTo(os.Stdout)
}
