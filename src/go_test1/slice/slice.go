package slice

import "fmt"

func NewSlice() []int  {
	slice := make([]int, 4)
	return  slice
}

func NewSlice1() []string  {
	slice := []string{"PHP","JAVA","GO"}
	return slice
}

func SliceAppend() []string{
	slice1 := make([]string,2);
	slice1[0] = "a"
	slice1[1] = "b"

	slice2 := []string{"c","d"}

	return append(slice1,slice2...)
}

func ShowSlice(slice []int)  {
	for _,k := range slice {
		fmt.Printf("key:%d\n",k)
	}
}

func ShowSlice1(slice []string)  {
	for _,k := range slice {
		fmt.Printf("key:%s\n",k)
	}
}
