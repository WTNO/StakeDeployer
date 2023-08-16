package step

import (
	"fmt"
	expect "github.com/google/goexpect"
	"github.com/wtno/StakeDeployer/command"
	"github.com/wtno/StakeDeployer/file"
	"os"
	"regexp"
	"strings"
	"time"
)

func Step1() {
	fmt.Println("Step 1 has started...")
	// 下载存款工具
	//command.RunCommand("/bin/bash", "-c", "curl -LO https://github.com/ethereum/staking-deposit-cli/releases/download/v2.5.0/staking_deposit-cli-d7b5304-linux-amd64.tar.gz")

	// 解压
	//command.RunCommand("/bin/bash", "-c", "tar xvf staking_deposit-cli-d7b5304-linux-amd64.tar.gz")

	// 运行存款工具
	e, _, err := expect.Spawn("./staking_deposit-cli-d7b5304-linux-amd64/deposit new-mnemonic --num_validators 2 --chain goerli --eth1_withdrawal_address 0x4D496CcC28058B1D74B7a19541663E21154f9c84", -1, expect.Verbose(true), expect.VerboseWriter(os.Stdout))
	fmt.Println("Spawn : " + e.String())
	if err != nil {
		fmt.Println(err)
	}
	defer e.Close()

	command.RunExpect(e, ".*Using the tool on an offline and secure device is highly recommended to keep your mnemonic safe..*", "\n")

	command.RunExpect(e, ".*Repeat your execution address for confirmation.*", "0x4D496CcC28058B1D74B7a19541663E21154f9c84\n")

	command.RunExpect(e, ".*Please choose the language of the mnemonic word list.*", "english\n")

	command.RunExpect(e, ".*Create a password that secures your validator keystore.*", "12345678\n")

	for {
		output, _, _ := command.RunExpect(e, ".*Repeat your keystore password for confirmation:.*", "12345678\n")
		if strings.Contains(output, "Repeat your keystore password") {
			break
		}
	}

	// 匹配并保存输出中的mnemonic
	mnemonic := matchMnemonic(e)

	// 输入上一步中的mnemonic
	typeMnemonic(e, mnemonic)

	re := regexp.MustCompile(".*Your keys can be found at.*")
	output, _, err := e.Expect(regexp.MustCompile("[a-zA-Z#]+"), 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	if re.MatchString(output) {
		e.Send("\n")
	}

	fmt.Println("Step 1 is over...")
}

// 匹配mnemonic
func matchMnemonic(e *expect.GExpect) string {
	re1 := regexp.MustCompile("\n\n.*\n\n")
	re2 := regexp.MustCompile(".*Press any key when you have written down your mnemonic.*")

	mnemonic := "mnemonic"
	for {
		output, _, err := e.Expect(regexp.MustCompile("[a-zA-Z]+"), 10*time.Second)

		if err != nil {
			fmt.Println(err)
		}

		if re1.MatchString(output) {
			mnemonic = strings.TrimSpace(re1.FindString(output))
			err = file.CreateAndWriteFile(mnemonic, "mnemonic.txt")
			if err != nil {
				fmt.Println(err)
			}
		}

		if re2.MatchString(output) {
			break
		}
	}
	e.Send("\n")
	return mnemonic
}

func typeMnemonic(e *expect.GExpect, mnemonic string) {
	re := regexp.MustCompile(".*type your mnemonic.*")
	for {
		output, _, err := e.Expect(regexp.MustCompile("[a-zA-Z#]+"), 10*time.Second)

		if err != nil {
			fmt.Println(err)
			e.Send("\n")
		}

		if re.MatchString(output) {
			e.Send(mnemonic + "\n")
			break
		}
	}
}

func Step2() {
	fmt.Println("Step 2 has started...")
	fmt.Println("连接到服务器:bu跳过")
	fmt.Println("Step 2 is over...")
}

func Step3() {
	fmt.Println("Step 3 has started...")

	fmt.Println("更新服务器：跳过")

	//command.RunCommand("/bin/bash", "-c", "sudo apt -y update && sudo apt -y upgrade")
	//command.RunCommand("/bin/bash", "-c", "sudo apt dist-upgrade && sudo apt autoremove")
	//command.RunCommand("/bin/bash", "-c", "sudo reboot")

	fmt.Println("Step 3 is over...")
}

func Step4() {
	fmt.Println("Step 4 has started...")

	fmt.Println("保护服务器：跳过")

	fmt.Println("Step 4 is over...")
}

func Step5() {
	fmt.Println("Step 5 has started...")

	fmt.Println("创建交换空间：跳过")

	fmt.Println("Step 5 is over...")
}

func Step6() {
	fmt.Println("Step 6 has started...")

	fmt.Println("配置时间同步：跳过")

	fmt.Println("Step 6 is over...")
}

// 生成客户端身份验证密钥
func Step7() {
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
