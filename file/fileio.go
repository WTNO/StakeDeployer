package file

import (
	"fmt"
	"io"
	"os"
)

func ReadAndWriteFile(inputFilePath, outputFilePath string) error {
	file, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("创建文件失败：", err)
		return err
	}

	// 读取文件内容
	//content, err := os.ReadFile(filePath)
	content, err := Asset(inputFilePath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return err
	}

	fmt.Println("content : ", string(content))

	_, err = io.WriteString(file, string(content))
	if err != nil {
		fmt.Println("写入文件失败：", err)
		return err
	}

	fmt.Println("文件写入成功")
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
