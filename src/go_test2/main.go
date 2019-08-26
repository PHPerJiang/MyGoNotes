package main

import (
	"go_test2/myinterface"
	"go_test2/ploymorphic"
)

func main() {
	//buffer.GetStdout()
	//接口
	//user := myinterface.User{"Jiang","jiang@163.com"}
	//myinterface.SendNotifier( user)
	//多态
	user := myinterface.User{"user","user@163.com"}
	amdin := ploymorphic.Admin{"admin","admin@163.com"}
	//使用指针接收者
	myinterface.SendNotifier(&user)
	//使用值接收者
	myinterface.SendNotifier(amdin)
}
