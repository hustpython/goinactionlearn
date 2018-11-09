// 这个示例程序展示如何序列化JSON字符串
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建一个保存键值对的映射
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "1",
		"cell": "2",
	}
	data, _ := json.MarshalIndent(c, "", "   ")
	fmt.Println(string(data))
}
