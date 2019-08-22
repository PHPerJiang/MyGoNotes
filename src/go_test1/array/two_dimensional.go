package array

import "fmt"

func NewArrayTwo() [3][3]int {
	var arr [3][3]int
	arr[0][1] = 1
	arr[0][2] = 2
	arr[1][1] = 3
	arr[1][2] =4
	return arr
}

func ShowArrayTwo(arr [3][3]int)  {
	for _,x := range arr{
		for _,y := range x{
			fmt.Println(y)
		}
	}
}
