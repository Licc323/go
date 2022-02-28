package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件
func readFromFile1()  {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//读文件
	//var tmp = make([]byte, 128) // 指定读的长度
	var tmp [128]byte
	for  {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Printf("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from filed, err:%v\n", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 0 {
			return
		}
	}

}

//利用bufio这个包读文件
func readFromFilebyBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//创建一个用来读从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line filed, err:%v", err)
			return
		}
		fmt.Print(line)
	}
}
func rwadFromFileByIoutil()  {
	//fileObj, err := os.Open("./main.go")
	//if err != nil {
	//	fmt.Printf("open file failed, err:%v", err)
	//	return
	//}
	////记得关闭文件
	//defer fileObj.Close()
	ret, err :=ioutil.ReadFile("./main.go")
	if err != nil{
		fmt.Printf("read file failed, err:%\n", err)
		return
	}
	fmt.Println(string(ret))
}
func main()  {
	// readFromFile1()
	//readFromFilebyBufio()
	rwadFromFileByIoutil()
}
