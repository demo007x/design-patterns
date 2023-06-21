package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
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
				file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}

				bf := bufio.NewReader(file)
				context, _, _ = bf.ReadLine()
				str := strings.Trim(string(context), "#")
				if deep > 2 {
					dir, fn := path.Split(filePath)
					dirSlice := strings.SplitAfterN(dir, "/", 3)
					return fmt.Sprintf("\n\r- [%s](%s)", str, dirSlice[len(dirSlice)-1]+fn)
				}
				file.Close()
				if deep == 2 {
					file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
					if err != nil {
						fmt.Println(err.Error())
						continue
					}
					write := bufio.NewWriter(file)
					write.WriteString(xxx)
					if err := write.Flush(); err != nil {
						fmt.Println(err.Error())
					}
					file.Close()
					// fmt.Println(xxx)
					// fmt.Println(filePath)
				}
			}
		}
	}
	return xxx
}
