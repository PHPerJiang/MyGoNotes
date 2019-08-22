package array

import (
	"fmt"
)

func NewArrayOne() [4]int{
	var arr [4]int
	arr[0] = 0
	arr[1] = 6
	arr[2] = 1
	arr[3] = 3
	return arr
}

func ShowArrayOne(arr [4]int)  {
	for x := range arr{
		fmt.Println(x)
	}
}
