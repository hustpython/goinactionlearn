// 这个示例程序展示如何解码JSON字符串到map变量中
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// JSON包含用于反序列化的演示字符串
var JSON = `{
	"name":"Gopher",
	"title":"programmer",
	"contact":{
		"home":"415.333.3335",
		"cell":"415.555.5555"
	}
}`

func main() {
	// 将JSON字符串反序列化到map变量中
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR", err)
		return
	}
	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact")
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])
}
