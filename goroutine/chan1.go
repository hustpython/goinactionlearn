package main

func main() {
	// 无缓冲的整型通道
	unbuffered := make(chan int)
	// 有缓冲的字符串通道
	buffered := make(chan string, 10)
	unbuffered <- 1
	buffered <- "Gopher"
	value := <-buffered
	valuestr := unbuffered
	//fmt.Println(value)
	//fmt.Println(valuestr)
}
