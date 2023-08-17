package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ReadAndWriteFile(inputFilePath, outputFilePath string) error {
	file, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("创建文件失败：", err)
		return err
	}

	// 获取当前可执行文件的绝对路径
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("获取可执行文件路径失败：", err)
		return err
	}

	// 获取可执行文件所在的目录路径
	dirPath := filepath.Dir(exePath)
	fmt.Println("dirPath", dirPath)

	// 构建要读取的文件的路径
	filePath := filepath.Join(dirPath, inputFilePath)
	fmt.Println("filePath", filePath)

	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return err
	}

	fmt.Println("content : ", content)

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
