// 这个示例程序展示如何解码JSON字符串到结构体中
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

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
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(c)
}
