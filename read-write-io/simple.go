package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	const fileUrl = `read-write-io/testwritefile.txt`
	Ioutil(fileUrl) //读
	content := "Hello!\n"
	WriteWithIoutil(fileUrl, content)
}

/**
读
*/
func Ioutil(name string) {
	if contents, err := ioutil.ReadFile(name); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		fmt.Println(`读取成功`, result)
	}
}

func WriteWithIoutil(name, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		fmt.Println("写入文件成功:", content)
	}
}
