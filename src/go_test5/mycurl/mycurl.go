package mycurl

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Mycurl()  {
	//抓取链接
	r,err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatalln(err)
	}
	//创建文件
	file,err := os.Create("webBody")
	if err != nil {
		log.Fatalln(err)
	}

	//关闭要创建的文件
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	//同时向标准输出及文件进行写操作
	dest := io.MultiWriter(os.Stdout, file)

	//读出响应内容并写入两个目的地
	_,err = io.Copy(dest,r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//关闭body
	if err := r.Body.Close(); err != nil {
		log.Fatalln(err)
	}

}
