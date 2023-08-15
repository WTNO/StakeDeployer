package file

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReadAndWriteFile(inputFilePath, outputFilePath string) error {
	// 读取文件内容
	content, err := os.ReadFile(inputFilePath)
	if err != nil {
		return err
	}

	// 写入文件
	err = os.WriteFile(outputFilePath, content, 0644)
	if err != nil {
		return err
	}

	log.Printf("write to %s success!", outputFilePath)
	return nil
}

func CreateAndWriteFile(content, outputFilePath string) error {
	file, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("创建文件失败：", err)
		return err
	}

	_, err = io.WriteString(file, content)
	if err != nil {
		fmt.Println("写入文件失败：", err)
		return err
	}

	fmt.Println("文件写入成功")
	return nil
}
