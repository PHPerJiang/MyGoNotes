package mychan

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func PlayBall(){
	//创建一个无缓冲通道
	court := make(chan int)

	//开启计数,2个routine
	wg.Add(2)

	//设置两个选手
	go player("JJJ",court)
	go player("ZZZ",court)

	//发球
	court <- 1

	//等待routine结束
	wg.Wait()
}

func player(name string, court chan int)  {
	//结束后调用done告诉wg工作完成
	defer wg.Done()

	for{
		//获取球
		ball, ok := <- court
		if !ok{
			fmt.Printf("Player %s Won!\n",name)
			return
		}

		//如果拿到球了,判断当前用户是否能击打回去
		n := rand.Intn(100)
		if n % 13 == 0 {
			fmt.Printf("Player %s Miss the ball!\n",name)
			close(court)
			return
		}

		//如果当前用户将球击回去了
		fmt.Printf("Player %s hit the ball %d!\n",name, ball)

		//击球次数+1
		ball++

		//将球打出去
		court <- ball
	}
}
