package command

import (
	"bytes"
	"fmt"
	expect "github.com/google/goexpect"
	"os/exec"
	"regexp"
	"strings"
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
	fmt.Println("RunCommand : ", name, arg)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

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
			if err.Error() != "EOF" {
				fmt.Println("Read Error : " + err.Error())
			}
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		fmt.Println("Wait Error : " + err.Error())
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}

	return nil
}

func RunExpect(e *expect.GExpect, regexpStr, sendStr string) (string, []string, error) {
	output, match, err := runExpect(e, regexpStr, sendStr)
	if err != nil {
		panic(err)
	}

	return output, match, err
}

func runExpect(e *expect.GExpect, regexpStr, sendStr string) (string, []string, error) {

	output, match, err := e.Expect(regexp.MustCompile(regexpStr), 10*time.Second)

	if err != nil {
		fmt.Println("Expect Error : ", err)
	}

	if len(strings.TrimSpace(output)) == 0 {
		fmt.Println("Not Match!")
	}

	err = e.Send(sendStr)
	if err != nil {
		fmt.Println("Send Error : ", err)
	}
	return output, match, nil
}

func RunExpectBatch(e *expect.GExpect, command string, dials ...Dial) error {
	return nil
}

func ExpectRun(cmd *exec.Cmd, expectStr, sendStr string) error {
	writer, _ := cmd.StdinPipe()
	reader, _ := cmd.StdoutPipe()

	tmp := make([]byte, 4096)
	_, err := reader.Read(tmp)
	if err != nil {
		fmt.Println("Read Error : ", err)
		return err
	}

	if !strings.Contains(string(tmp), expectStr) {
		fmt.Println("Not Match ~")
	}

	_, err = writer.Write([]byte(sendStr))
	if err != nil {
		fmt.Println("Write Error : ", err)
		return err
	}

	return nil
}
