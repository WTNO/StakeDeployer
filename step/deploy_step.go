package step

import (
	"fmt"
	"sync"
)

type MyStruct struct {
	S string
}

var StepMap = map[string]interface{}{
	"step1": Step1,
}

func (ms MyStruct) Print() string {
	return ms.S
}

func Step1(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	// 第一个步骤的逻辑
	fmt.Println("Step 1")
	// 完成第一个步骤后，向通道发送信号
	ch <- true
}

func Step2(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	// 等待第一个步骤完成的信号
	<-ch
	// 第二个步骤的逻辑
	fmt.Println("Step 2")
	// 完成第二个步骤后，向通道发送信号
	ch <- true
}

func Step3(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	// 等待第二个步骤完成的信号
	<-ch
	// 第三个步骤的逻辑
	fmt.Println("Step 3")
	// 完成第三个步骤后，向通道发送信号
	ch <- true
}

func Step4(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	// 等待第三个步骤完成的信号
	<-ch
	// 第四个步骤的逻辑
	fmt.Println("Step 4")
	// 完成第四个步骤后，向通道发送信号
	ch <- true
}

func Step5(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	// 等待第四个步骤完成的信号
	<-ch
	// 第五个步骤的逻辑
	fmt.Println("Step 5")
	close(ch)
}
