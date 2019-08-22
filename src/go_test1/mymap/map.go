package mymap

import "fmt"

func CreateMap1() map[string]int {
	mymap := make(map[string]int)
	mymap["PHP"] = 1
	mymap["JAVA"] =2
	return mymap
}

func CreateMap2() map[string]int  {
	mymap := map[string]int{"PHP":1,"JAVA":2}
	return mymap
}

func ShowMap(mymap map[string]int)  {
	for k,v := range mymap {
		fmt.Printf("key:%s,value:%d \n",k,v)
	}
}