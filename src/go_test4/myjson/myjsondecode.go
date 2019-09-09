package myjson

import (
	"encoding/json"
	"fmt"
	"log"
)

type Content struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Contact struct{
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

var myjson = `{
	"name":"Gopher",
	"title":"this is test",
	"contact":{
		"home":"qingdao",
		"cell":"123"
	}
}`

func JsonDecode2Var()  {
	var C Content
	err := json.Unmarshal([]byte(myjson),&C)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(C)
}

func JsonDecode2Map()  {
	M := make(map[string] interface{})
	err := json.Unmarshal([]byte(myjson),&M)
	if err != nil{
		panic(err)
		return
	}
	for k,v := range M{
		fmt.Printf("key: %s => value: %s\n",k,v)
	}
}

