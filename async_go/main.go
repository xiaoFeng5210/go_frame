package main

import (
	"fmt"
	"time"
)

func doSome(number int) {
	time.Sleep(time.Second * time.Duration(number))
	fmt.Println("任务完成", number)
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		doSome(i)
	}
	fmt.Println("程序运行时间", time.Since(start))
}
