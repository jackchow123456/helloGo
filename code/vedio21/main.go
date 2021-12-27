package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile() {
	// 打开文件
	fileObj, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	// 关闭文件
	defer fileObj.Close()

	// 读取文件
	buf := make([]byte, 128)
	n, err := fileObj.Read(buf)
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	fmt.Printf("read %d from file.\n", n)
	fmt.Println(string(buf))
}

func readAll() {
	fileObj, err := os.Open("./xxx.txt")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	// 关闭文件
	defer fileObj.Close()

	// 读取文件
	for {
		buf := make([]byte, 128)
		n, err := fileObj.Read(buf)
		if err == io.EOF {
			fmt.Println("hhh")
			fmt.Println(string(buf[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v\n", err)
			return
		}
		fmt.Printf("read %d byte from file.\n", n)
		fmt.Println(string(buf[:n]))
	}
}

func readByIoutil() {
	content, err := ioutil.ReadFile("./xxx.txt")
	if err != nil {
		fmt.Printf("read file by ioutil failed,err:%v\n", err)
		return
	}
	fmt.Println(string(content))
}

func write() {
	fileObj, err := os.OpenFile("./xx23.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()

	str := "小明小不小"
	fileObj.Write([]byte(str))
	fileObj.WriteString("hello man")
}

// 文件操作

func main() {
	// readFile()
	// readAll()
	// readByIoutil()
	write()
}
