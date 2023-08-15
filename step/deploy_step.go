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

	fmt.Println("mnemonic start")

	// 这一步需要保存mnemonic
	mnemonic := matchMnemonic(e)

	// TODO：问题停留在这一步
	// 输入上一步中的mnemonic
	//for {
	//	output, _, _ := command.RunExpect(e, ".*Please type your mnemonic (separated by spaces) to confirm you have written it down.*", mnemonic+"\n")
	//	fmt.Println("****************", output)
	//	if strings.Contains(output, "Please type your mnemonic") {
	//		break
	//	}
	//}
	typeMnemonic(e, mnemonic)

	command.RunExpect(e, ".*Your keys can be found at.*", "\n")

	fmt.Println("Step 1 is over...")
}

// 匹配mnemonic
func matchMnemonic(e *expect.GExpect) string {
	re1 := regexp.MustCompile("\n\n.*\n\n")
	re2 := regexp.MustCompile(".*Press any key when you have written down your mnemonic.*")

	mnemonic := "mnemonic"
	for {
		output, _, err := e.Expect(regexp.MustCompile("[a-zA-Z]+"), 10*time.Second)
		fmt.Println(output)

		if err != nil {
			fmt.Println(err)
		}

		if re1.MatchString(output) {
			mnemonic = strings.TrimSpace(re1.FindString(output))
			fmt.Println("mnemonic : ", mnemonic)
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
	//re := regexp.MustCompile(".*type your mnemonic.*")
	for {
		_, _, err := e.Expect(regexp.MustCompile("[a-zA-Z]+"), 10*time.Second)

		if err != nil {
			fmt.Println(err)
		}

		e.Send(mnemonic + "\n")
	}
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
