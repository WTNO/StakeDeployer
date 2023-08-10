package main

import (
	"flag"
	"fmt"
	"github.com/wtno/StakeDeployer/step"
)

func main() {
	// 解析命令行参数
	stepNum := flag.Int("step", 1, "the step to execute")
	flag.Parse()

	// 根据命令行参数决定执行哪个步骤的函数
	switch *stepNum {
	case 1:
		step.Step1()
		fallthrough
	case 2:
		step.Step2()
		fallthrough
	case 3:
		step.Step3()
		fallthrough
	case 4:
		step.Step4()
		fallthrough
	case 5:
		step.Step5()
		fallthrough
	case 6:
		step.Step6()
		fallthrough
	case 7:
		step.Step7()
		fallthrough
	case 8:
		step.Step8()
		fallthrough
	case 9:
		step.Step9()
		fallthrough
	case 10:
		step.Step10()
		fallthrough
	case 11:
		step.Step11()
		fallthrough
	case 12:
		step.Step12()
		fallthrough
	case 13:
		step.Step13()
		fallthrough
	case 14:
		step.Step14()
		fallthrough
	case 15:
		step.Step15()
		fallthrough
	case 16:
		step.Step16()
		fallthrough
	case 17:
		step.Step17()
		fallthrough
	case 18:
		step.Step18()
		fallthrough
	default:
		fmt.Println("The deployment execution has concluded.")
	}
}
