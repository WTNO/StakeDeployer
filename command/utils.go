package command

import (
	"fmt"
	expect "github.com/google/goexpect"
	"os/exec"
	"regexp"
	"time"
)

type Dial struct {
	expect string
	send   string
}

func RunCommand(name string, arg ...string) {
	err := runCommand(name, arg...)
	if err != nil {
		panic(err)
	}
}

func runCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...) // 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		fmt.Println("StdoutPipe Error : " + err.Error())
		return err
	}

	if err = cmd.Start(); err != nil {
		fmt.Println("Start Error : " + err.Error())
		return err
	}

	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			fmt.Println("Read Error : " + err.Error())
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		fmt.Println("Wait Error : " + err.Error())
		return err
	}

	return nil
}

func RunExpect(e *expect.GExpect, regexpStr, sendStr string) error {
	output, _, err := e.Expect(regexp.MustCompile(regexpStr), 10*time.Minute)
	if err != nil {
		fmt.Println("RunExpect Error : ", err)
	}

	fmt.Println(output)

	err = e.Send(sendStr)
	if err != nil {
		fmt.Println("RunExpect Error : ", err)
	}
	return nil
}

func RunExpectBatch(e *expect.GExpect, command string, dials ...Dial) error {
	return nil
}