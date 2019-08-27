package mygoroutine

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func RoutineTest(){
	//分配一个逻辑处理器
	runtime.GOMAXPROCS(1)
	//增加等待两个routine
	wg.Add(2)

	//开启两个routine
	println("Create two goroutine...")
	go printPrime("A")
	go printPrime("B")

	//等待routine结束
	fmt.Println("Wait routine finish")
	wg.Wait()

	println("Terminating Program")
}


//显示5000内的素数
func printPrime(prefix string)  {
	defer wg.Done()
next:
	for i:= 2; i<10000; i++  {
		for j:=2; j<i; j++ {
			if i%j == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix,i)
	}
	println("Completed",prefix)
}
