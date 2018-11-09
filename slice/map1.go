//映射是一个储存键值对的无序组合
package main 

import (
	"fmt"
)
func removeColor(colordd map[string]string,key string){
	delete(colordd,key)
}
func main(){
	//创建一个映射,键的类型是strig,值是int
	dict:=make(map[string]int)
	//创建一个映射,键和值的类型都是string
	dictstr:=map[string]string{"red":"红","Orange":"橙"}
	fmt.Println(dict)
	fmt.Println(dictstr)

	colors := map[string]string{}
	colors["blue"] = "绿"

	value,exists := colors["blue"]
	fmt.Println(colors["blue"])
	if exists{
		fmt.Println(value)
	}
	//迭代
	colorsdd := map[string]string{
		"AliceBlue":   "#f0f8ff",  "Coral":       "#ff7F50",  "DarkGray":    "#a9a9a9",  "ForestGreen":"#228b22", 
	}
	// 显示映射里的所有颜色
	for key, value := range colorsdd {  
		fmt.Printf("Key: %s  Value: %s\n", key, value) 
	}
	//删除
	//delete(colorsdd,"Coral")
	removeColor(colorsdd,"Coral")
	for key, value := range colorsdd {  
		fmt.Printf("afterdeletedkey: %s  Value: %s\n", key, value) 
	}
}