package main

import (
	"fmt"
	"go_test1/mystruct"
)


func main()  {
	fmt.Println("hello word")
	//array.ShowArrayOne(array.NewArrayOne())
	//array.ShowArrayTwo(array.NewArrayTwo())
	//slice.ShowSlice(slice.NewSlice())
	//slice.ShowSlice1(append(slice.NewSlice1(), "NODE"))
	//slice.ShowSlice1(slice.SliceAppend());
	//fmt.Println(len(slice.SliceAppend()))
	//mymap.ShowMap(mymap.CreateMap2())
	mystruct.ShowStruct(mystruct.CreateStruct())
}




