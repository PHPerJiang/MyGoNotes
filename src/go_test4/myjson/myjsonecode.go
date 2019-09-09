package myjson

import (
	"encoding/json"
	"fmt"
	"log"
)

func Map2Json()  {
	M := make(map[string]interface{})

	M["name"] = "Gopher"
	M["age"]  = 23
	M["sex"]  = "man"

	//为要转码的map增加前缀及缩进
	data, err := json.MarshalIndent(M,"json:","   ")
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(string(data))
}

func Map2JSON1(){
	M := make(map[string]interface{})

	M["name"] = "Gopher"
	M["age"]  = 23
	M["sex"]  = "man"

	//不加前缀及缩进的转码方式，常用于web-api
	data,err := json.Marshal(M)

	if err != nil {
		log.Println(err)
		return
	}

	println(string(data))
}
