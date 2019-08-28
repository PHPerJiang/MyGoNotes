package mychan

import (
	"fmt"
	//"sync"
	"time"
)

//var wg sync.WaitGroup

func Relay()  {

	//创建无缓冲通道
	baton := make(chan int)

	//增加计数
	wg.Add(1)

	//给第一位跑者接力棒
	go Runner(baton)

	//开始比赛
	baton <-1

	//等待跑完
	wg.Wait()
}

func Runner(baton chan int){
	//默认结束人物时计数减一
	//defer wg.Done()

	//声明下一棒跑者
	var newRunner int

	//获取接力棒
	runner := <-baton

	//拿到接力棒开始跑
	fmt.Printf("Runner %d running with baton\n", runner)

	//如果不是最后一棒则创建下一位接力者进入跑道
	if runner != 4 {
		newRunner = runner+1
		fmt.Printf("Runner %d to the line\n",newRunner)
		//开跑，进入阻塞模式
		go Runner(baton)
	}

	//当前跑者围着跑道跑
	time.Sleep(time.Microsecond * 100)

	//如果是最后一位跑者则比赛结束
	if runner == 4 {
		fmt.Printf("Runner %d finished, race over\n",runner)
		//任务结束，计量数减一
		wg.Done()
		close(baton)
		return
	}

	//将接力棒交给下一位跑者
	fmt.Printf("Runner %d exchange with runner %d\n",runner,newRunner)

	//传递接力棒
	baton <- newRunner
}
