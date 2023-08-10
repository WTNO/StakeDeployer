package step

import (
	"fmt"
	"github.com/wtno/StakeDeployer/command"
)

func Step1() {
	fmt.Println("Step 1 has started...")
	// 下载存款工具
	//command.RunCommand("/bin/bash", "-c", "curl -LO https://github.com/ethereum/staking-deposit-cli/releases/download/v2.5.0/staking_deposit-cli-d7b5304-linux-amd64.tar.gz")

	// 解压
	//command.RunCommand("/bin/bash", "-c", "tar xvf staking_deposit-cli-d7b5304-linux-amd64.tar.gz")
	//command.RunCommand("/bin/bash", "-c", "cd staking_deposit-cli-d7b5304-linux-amd64")
	//command.RunCommand("/bin/bash", "-c", "ls")

	// 运行存款工具
	//err := command.RunCommand("/bin/bash", "-c", "cd staking_deposit-cli-d7b5304-linux-amd64 && ./deposit new-mnemonic --num_validators 2 --chain goerli --eth1_withdrawal_address 0x4D496CcC28058B1D74B7a19541663E21154f9c84")
	err := command.RunCommand("sh", "-c", "cd staking_deposit-cli-d7b5304-linux-amd64 && ls")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Step 1 is over...")
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
