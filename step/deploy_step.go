package step

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func Step1() {
	cmd := exec.Command("/bin/bash", "-c", "curl -LO https://github.com/ethereum/staking-deposit-cli/releases/download/v2.5.0/staking_deposit-cli-d7b5304-linux-amd64.tar.gz")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("error:\n%s\n", stderr.String())
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}

	fmt.Printf("output:\n%s\nerror:\n%s\n", stdout.String(), stderr.String())
}

func Step2() {
	// 第二个步骤的逻辑
	fmt.Println("Step 2")
}

func Step3() {
	// 第三个步骤的逻辑
	fmt.Println("Step 3")
}

func Step4() {
	// 第四个步骤的逻辑
	fmt.Println("Step 4")
}

func Step5() {
	// 第五个步骤的逻辑
	fmt.Println("Step 5")
}

func Step6() {
	// 第五个步骤的逻辑
	fmt.Println("Step 6")
}

func Step7() {
	// 第五个步骤的逻辑
	fmt.Println("Step 7")
}

func Step8() {
	// 第五个步骤的逻辑
	fmt.Println("Step 8")
}

func Step9() {
	// 第五个步骤的逻辑
	fmt.Println("Step 9")
}

func Step10() {
	// 第五个步骤的逻辑
	fmt.Println("Step 10")
}

func Step11() {
	// 第五个步骤的逻辑
	fmt.Println("Step 11")
}

func Step12() {
	// 第五个步骤的逻辑
	fmt.Println("Step 12")
}

func Step13() {
	// 第五个步骤的逻辑
	fmt.Println("Step 13")
}

func Step14() {
	// 第五个步骤的逻辑
	fmt.Println("Step 14")
}

func Step15() {
	// 第五个步骤的逻辑
	fmt.Println("Step 15")
}

func Step16() {
	// 第五个步骤的逻辑
	fmt.Println("Step 16")
}

func Step17() {
	// 第五个步骤的逻辑
	fmt.Println("Step 17")
}

func Step18() {
	// 第五个步骤的逻辑
	fmt.Println("Step 18")
}
