package embedding

import "fmt"

type Notifier interface {
	Notify()
}

type User struct {
	Name string
	Email string
}

type Admin struct {
	User
	Level int
}

func (u *User)Notify()  {
	fmt.Printf("Send user email to %s<%s>\n",u.Name,u.Email)
}

//func (a *Admin)Notify()  {
//	fmt.Printf("Send user email to %s<%s>\n,and this user level is %d\n",a.User.Name,a.User.Email,a.Level)
//}

func SendNotifier(n Notifier)  {
	n.Notify()
}
