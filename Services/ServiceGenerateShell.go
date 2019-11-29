package Services

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)



func GenerateShell(fileName string, content string)  {
	path := "./Shell/"+fileName+".sh"
	_ = os.Remove(path)
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
		//func NewWriter(w io.Writer) *Writer
		bufferWrite := bufio.NewWriter(f)
	if err != nil {
		panic(err)
	}
	contentShell := "#!/bin/sh\r"+content
	for _, v := range contentShell {
		//将数据写入缓冲区
		//func (b *Writer) WriteString(s string) (int, error)
		bufferWrite.WriteString(string(v))
	}
	data, _ := ioutil.ReadFile(path)
	//fmt.Println(string(data))   //空的内容

	//将缓冲区的数据写入文件
	//func (b *Writer) Flush() error
	bufferWrite.Flush()

	data, _ = ioutil.ReadFile(path)
	fmt.Println(string(data))   //1234567890
}