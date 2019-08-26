package myinterface

import "fmt"

type Notifier interface {
	Notify()
}

type User struct {
	Name string
	Email string
}
//使用指针接收者只能解析指针类型
func (u *User)Notify()  {
	fmt.Printf("Send user email to %s<%s>\n",u.Name,u.Email)
}
//使用值接收者能接受指针类型或者值类型，如果接受的指针类型的go自动解引用
func SendNotifier(n Notifier)  {
	n.Notify()
}