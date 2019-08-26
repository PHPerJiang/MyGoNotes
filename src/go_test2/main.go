package main

import (
	"go_test2/embedding"
)

func main() {
	//buffer.GetStdout()

	//接口
	//user := myinterface.User{"Jiang","jiang@163.com"}
	//myinterface.SendNotifier( user)

	//多态
	//user := myinterface.User{"user","user@163.com"}
	//amdin := ploymorphic.Admin{"admin","admin@163.com"}
	//使用指针接收者
	//myinterface.SendNotifier(&user)
	//使用值接收者
	//myinterface.SendNotifier(amdin)

	//嵌入类型
	//embedding中并没有实现notifier的admin接口，但因为admin中镶嵌了user类型，user类型实现了notifier接口，
	//当传递admin指针给输出方法时，admin也间接实现了notifier接口
	ad := embedding.Admin{
		User:embedding.User{"user","user@163.com"},
		Level:1,
	}
	embedding.SendNotifier(&ad)
}
