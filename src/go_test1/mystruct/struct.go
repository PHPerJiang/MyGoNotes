package mystruct

import "fmt"

type User struct {
	name string
	age  int
}

func CreateStruct() User {
	user := User{name:"PHPer",age:23}
	return user
}

func ShowStruct(user User){
	fmt.Printf("name: %s, age:%d",user.name,user.age)
}