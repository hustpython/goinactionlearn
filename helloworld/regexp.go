package main 

import (
	"fmt"
	"regexp"
)

func main(){
	flysnowRegexp := regexp.MustCompile(`^http://www.flysnow.org/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`)
	params := flysnowRegexp.FindStringSubmatch("http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html")
	for _,param := range params{
		fmt.Println(param)
	}
}