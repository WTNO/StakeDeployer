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
	command.RunCommand("/bin/bash", "-c", "tar xvf staking_deposit-cli-d7b5304-linux-amd64.tar.gz")

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
		time.Sleep(10 * time.Second)
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
	fmt.Println("连接到服务器:跳过")
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
	fmt.Println("Step 7 has started...")

	command.RunSudoCommand("/bin/bash", "-c", "sudo mkdir -p /var/lib/jwtsecret")
	command.RunCommand("/bin/bash", "-c", "openssl rand -hex 32 | sudo tee /var/lib/jwtsecret/jwt.hex > /dev/null")

	fmt.Println("Step 7 is over...")
}

// 配置执行客户端Geth
func Step8() {
	fmt.Println("Step 8 has started...")

	// 下载
	//command.RunCommand("/bin/bash", "-c", "curl -LO https://gethstore.blob.core.windows.net/builds/geth-linux-amd64-1.12.2-bed84606.tar.gz")

	// 解压
	command.RunCommand("/bin/bash", "-c", "tar xvf geth-linux-amd64-1.12.2-bed84606.tar.gz")

	// 复制
	command.RunSudoCommand("/bin/bash", "-c", "sudo cp ~/geth-linux-amd64-1.12.2-bed84606/geth /usr/local/bin")

	// 删除压缩包和解压文件夹
	command.RunCommand("/bin/bash", "-c", "rm geth-linux-amd64-1.12.2-bed84606.tar.gz")
	command.RunCommand("/bin/bash", "-c", "rm -r geth-linux-amd64-1.12.2-bed84606")

	// 为服务创建一个账户以运行
	command.RunSudoCommand("/bin/bash", "-c", "sudo useradd --no-create-home --shell /bin/false geth")

	// 创建数据目录。这是存储以太坊区块链数据所必需的。
	command.RunSudoCommand("/bin/bash", "-c", "sudo mkdir -p /var/lib/geth")

	// 设置目录权限。geth用户账户需要有修改数据目录的权限。
	command.RunSudoCommand("/bin/bash", "-c", "sudo chown -R geth:geth /var/lib/geth")

	// 创建一个systemd服务配置文件来配置服务。
	file.ReadAndWriteFile("config/geth.service", "/etc/systemd/system/geth.service")

	// 重新加载systemd以反映更改并启动服务。检查状态以确保它正常运行。
	command.RunSudoCommand("/bin/bash", "-c", "sudo systemctl daemon-reload")
	command.RunSudoCommand("/bin/bash", "-c", "sudo systemctl start geth")
	command.CheckServiceRunning("geth")

	// 启用geth服务以在重新启动时自动启动。
	command.RunSudoCommand("/bin/bash", "-c", "sudo systemctl enable geth")

	fmt.Println("Step 8 is over...")
}

// 安装Prysm共识客户端
func Step9() {
	fmt.Println("Step 9 has started...")

	// 下载Prysm共识客户端
	//command.RunCommand("/bin/bash", "-c", "curl -LO https://github.com/prysmaticlabs/prysm/releases/download/v4.0.7/beacon-chain-v4.0.7-linux-amd64")
	//command.RunCommand("/bin/bash", "-c", "curl -LO https://github.com/prysmaticlabs/prysm/releases/download/v4.0.7/validator-v4.0.7-linux-amd64")

	// 重命名文件并使其可执行
	command.RunCommand("/bin/bash", "-c", "mv beacon-chain-v4.0.7-linux-amd64 beacon-chain")
	command.RunCommand("/bin/bash", "-c", "mv validator-v4.0.7-linux-amd64 validator")
	command.RunCommand("/bin/bash", "-c", "chmod +x beacon-chain")
	command.RunCommand("/bin/bash", "-c", "chmod +x validator")
	command.RunSudoCommand("/bin/bash", "-c", "sudo cp beacon-chain /usr/local/bin")
	command.RunSudoCommand("/bin/bash", "-c", "sudo cp validator /usr/local/bin")

	// 清理文件夹
	command.RunCommand("/bin/bash", "-c", "rm beacon-chain && rm validator")

	fmt.Println("Step 9 is over...")
}

// 导入验证者密钥
func Step10() {
	fmt.Println("Step 10 has started...")

	// 将验证器密钥库文件导入Prysm
	command.RunSudoCommand("/bin/bash", "-c", "sudo mkdir -p /var/lib/prysm/validator")
	command.RunSudoCommand("/bin/bash", "-c", "sudo chown -R root:root /var/lib/prysm/validator")

	// 下面这一步有互动过程
	e, _, err := expect.Spawn("/usr/local/bin/validator accounts import --keys-dir=$HOME/validator_keys --wallet-dir=/var/lib/prysm/validator --goerli", -1, expect.Verbose(true), expect.VerboseWriter(os.Stdout))
	fmt.Println("Spawn : " + e.String())
	if err != nil {
		fmt.Println(err)
	}
	defer e.Close()

	// 创建钱包密码
	command.RunExpect(e, ".*to accept this terms and conditions.*", "accept\n")
	command.RunExpect(e, ".*New wallet password.*", "cptbtptp\n")
	command.RunExpect(e, ".*Confirm password.*", "cptbtptp\n")

	// 这一步是在下面代码出现报错expect: Process not running时执行的
	//command.RunExpect(e, ".*Wallet password.*", "cptbtptp\n")

	// 输入第一步中创建密钥时提供的密码
	command.RunExpect(e, ".*Enter the password for your.*", "12345678\n")
	command.RunExpect(e, ".*Importing accounts.*", "")

	// 创建钱包密码文件
	file.CreateAndWriteFile("cptbtptp", "/var/lib/prysm/validator/password.txt")

	fmt.Println("Step 10 is over...")
}

// 配置Beacon节点服务
func Step11() {
	fmt.Println("Step 11 has started...")

	// 设置帐户
	command.RunSudoCommand("/bin/bash", "-c", "sudo useradd --no-create-home --shell /bin/false prysmbeacon")

	// 设置目录和权限
	command.RunSudoCommand("/bin/bash", "-c", "sudo mkdir -p /var/lib/prysm/beacon")
	command.RunSudoCommand("/bin/bash", "-c", "sudo chown -R prysmbeacon:prysmbeacon /var/lib/prysm/beacon")

	// 创建并配置服务
	file.ReadAndWriteFile("config/prysmbeacon.service", "/etc/systemd/system/prysmbeacon.service")

	// 重新加载systemd以反映更改并启动服务。检查状态以确保它正常运行。
	command.RunSudoCommand("/bin/bash", "-c", "sudo systemctl daemon-reload")
	command.RunSudoCommand("/bin/bash", "-c", "sudo systemctl start prysmbeacon")
	command.CheckServiceRunning("prysmbeacon")

	// 启用geth服务以在重新启动时自动启动。
	command.RunSudoCommand("/bin/bash", "-c", "sudo systemctl enable prysmbeacon")

	fmt.Println("Step 11 is over...")
}

// 配置验证器服务
func Step12() {
	fmt.Println("Step 12 has started...")

	// 设置帐户
	command.RunSudoCommand("/bin/bash", "-c", "sudo useradd --no-create-home --shell /bin/false prysmvalidator")

	// 在第10步中，验证器导入过程创建了以下目录：/var/lib/prysm/validator。设置目录权限，以便prysmvalidator账户可以修改该目录。
	command.RunSudoCommand("/bin/bash", "-c", "sudo chown -R prysmvalidator:prysmvalidator /var/lib/prysm/validator")

	// 创建并配置服务
	file.ReadAndWriteFile("config/prysmvalidator.service", "/etc/systemd/system/prysmvalidator.service")

	// 重新加载systemd以反映更改并启动服务。检查状态以确保它正常运行。
	command.RunSudoCommand("/bin/bash", "-c", "sudo systemctl daemon-reload")
	command.RunSudoCommand("/bin/bash", "-c", "sudo systemctl start prysmvalidator")
	command.CheckServiceRunning("prysmvalidator")

	// 启用geth服务以在重新启动时自动启动。
	command.RunSudoCommand("/bin/bash", "-c", "sudo systemctl enable prysmvalidator")

	fmt.Println("Step 12 is over...")
}

// 修复步骤10的问题
func Step13() {
	fmt.Println("Step 13 has started...")

	// 重新调用一次步骤1
	Step1()

	// 下面这一步有互动过程
	//e, _, err := expect.Spawn("/usr/local/bin/validator accounts import --keys-dir=$HOME/validator_keys --wallet-dir=/var/lib/prysm/validator --goerli", -1, expect.Verbose(true), expect.VerboseWriter(os.Stdout))
	//fmt.Println("Spawn : " + e.String())
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer e.Close()
	//
	//// 这一步是在下面代码出现报错expect: Process not running时执行的
	//command.RunExpect(e, ".*Wallet password.*", "cptbtptp\n")
	//
	//// 输入第一步中创建密钥时提供的密码
	//command.RunExpect(e, ".*Enter the password for your.*", "12345678\n")
	//command.RunExpect(e, ".*Importing accounts.*", "")
	//
	//// 创建钱包密码文件
	//file.CreateAndWriteFile("cptbtptp", "/var/lib/prysm/validator/password.txt")

	fmt.Println("Step 13 is over...")
}
