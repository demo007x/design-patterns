package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var root = "."
var deep = 1

func main() {
	visitDir(root, deep)
}

func visitDir(dir string, deep int) string {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		panic(err.Error())
	}
	var xxx string
	for _, entry := range dirEntries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		filePath := dir + "/" + info.Name()
		var context []byte
		if info.IsDir() {
			xxx = visitDir(filePath, deep+1)
		} else {
			if strings.Index(info.Name(), ".md") > 0 {
				// 读取 markdown 的文件的第一行
				file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_APPEND, 0666)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				defer file.Close()
				bf := bufio.NewReader(file)
				context, _, _ = bf.ReadLine()
				str := strings.Trim(string(context), "#")
				if deep > 2 {
					return fmt.Sprintf("\n\r- [%s](%s)", str, filePath)
				}
				if deep == 2 {
					write := bufio.NewWriter(file)
					write.WriteString(xxx)
					write.Flush()
				}
			}
		}
	}
	return xxx
}
