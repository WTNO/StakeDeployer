package step

import (
	"fmt"
	expect "github.com/google/goexpect"
	"github.com/wtno/StakeDeployer/command"
	"github.com/wtno/StakeDeployer/file"
	"regexp"
	"strings"
)

func Step1() {
	fmt.Println("Step 1 has started...")
	// 下载存款工具
	//command.RunCommand("/bin/bash", "-c", "curl -LO https://github.com/ethereum/staking-deposit-cli/releases/download/v2.5.0/staking_deposit-cli-d7b5304-linux-amd64.tar.gz")

	// 解压
	//command.RunCommand("/bin/bash", "-c", "tar xvf staking_deposit-cli-d7b5304-linux-amd64.tar.gz")

	// 运行存款工具
	e, _, err := expect.Spawn("./staking_deposit-cli-d7b5304-linux-amd64/deposit new-mnemonic --num_validators 2 --chain goerli --eth1_withdrawal_address 0x4D496CcC28058B1D74B7a19541663E21154f9c84", -1)
	fmt.Println("Spawn : " + e.String())
	if err != nil {
		fmt.Println(err)
	}
	defer e.Close()

	command.RunExpect(e, ".*Using the tool on an offline and secure device is highly recommended to keep your mnemonic safe..*", "\n")

	command.RunExpect(e, ".*Repeat your execution address for confirmation.*", "0x4D496CcC28058B1D74B7a19541663E21154f9c84\n")

	command.RunExpect(e, ".*Please choose the language of the mnemonic word list.*", "english\n")

	command.RunExpect(e, ".*Create a password that secures your validator keystore.*", "123456789\n")

	command.RunExpect(e, ".*Eepeat your keystore password for confirmatio.*", "123456789\n")

	// 这一步需要保存mnemonic
	// This is your mnemonic (seed phrase). Write it down and store it safely. It is the ONLY way to retrieve your deposit.
	//
	//
	// alarm fog wedding immense success couch staff future endless dilemma broccoli clever rotate humor endless toward laugh include loan invite segment use moral float
	//
	//
	// Press any key when you have written down your mnemonic.
	output, _, _ := command.RunExpect(e, ".*This is your mnemonic (seed phrase).*", "\n")
	re, _ := regexp.Compile("\n\n.*\n\n")
	mnemonic := strings.TrimSpace(re.FindString(output))
	file.WriteFile(mnemonic, "~/mnemonic")

	// 输入上一步中的mnemonic
	command.RunExpect(e, ".*Please type your mnemonic (separated by spaces) to confirm you have written it down.*", mnemonic)

	command.RunExpect(e, ".*Success!\nYour keys can be found at.*", "\n")

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
