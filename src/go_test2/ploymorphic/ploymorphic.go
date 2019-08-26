package ploymorphic

import "fmt"

type Notifier interface {
	Notify()
}

type User struct {
	Name string
	Email string
}

type Admin struct {
	Name string
	Email string
}

func (u User)Notify()  {
	fmt.Printf("Send user email to %s<%s>\n",u.Name,u.Email)
}

func (a Admin)Notify()  {
	fmt.Printf("Send user email to %s<%s>\n",a.Name,a.Email)
}

func sendNotifier(n Notifier)  {
	n.Notify()
}
