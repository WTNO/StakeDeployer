package file

import (
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

	log.Println("文件写入成功！")
	return nil
}
