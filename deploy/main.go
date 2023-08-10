package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	// 解析命令行参数
	step := flag.Int("step", 1, "the step to execute")
	flag.Parse()

	// 创建一个等待组和一个通道
	var wg sync.WaitGroup
	ch := make(chan bool)

	// 增加等待组的计数器
	wg.Add(5)

	// 根据命令行参数决定执行哪个步骤的函数
	switch *step {
	case 1:
		go step.Step1(&wg, ch)
	case 2:
		go step2(&wg, ch)
	case 3:
		go step3(&wg, ch)
	case 4:
		go step4(&wg, ch)
	case 5:
		go step5(&wg, ch)
	default:
		fmt.Println("Invalid step")
	}

	// 等待所有goroutine执行完毕
	wg.Wait()
}
