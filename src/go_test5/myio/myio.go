package myio

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func Io1()  {
	//创建一个buffer值，将数据写入
	var b bytes.Buffer
	b.Write([]byte("Hello "))
	//将数据追加到buffer里
	_,err := fmt.Fprintf(&b,"World")
	if err != nil {
		log.Println(err)
	}

	//数据输出到标准输出设备
	_,err = b.WriteTo(os.Stdout)
	if err != nil {
		log.Println(err)
	}
}
