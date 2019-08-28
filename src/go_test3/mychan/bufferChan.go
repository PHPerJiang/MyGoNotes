package mychan

import (
	"fmt"
	"time"
)

const (
	routineNum = 4    //协程数
	taskNum    = 10   //任务数
)

func BufferChan(){
	//创建任务通道
	taskChan := make(chan string,taskNum)

	//启动协程计量数
	wg.Add(routineNum)

	//启动协程
	for i:= 0 ; i< routineNum; i++ {
		go worker(taskChan, i+1)
	}

	//向任务通道打入任务
	for i:=0 ;i< taskNum ;i++  {
		taskChan <- fmt.Sprintf("Task : %d",i+1)
	}

	//任务进入通道后关闭通道
	close(taskChan)

	//等待协程完毕
	wg.Wait()
}

func worker(taskChan chan string,routineNum int){
	//任务结束 计量数减一
	defer wg.Done()
	for true {
		//获取通道任务
		task,ok := <-taskChan
		if !ok {
			fmt.Printf("Routine %d shutting down\n",routineNum)
			return
		}

		//如果获取到任务，则输出任务
		fmt.Printf("Routine %d get work %s , it is working\n",routineNum,task)
		//等待任务结束，事件延迟
		time.Sleep(time.Second * 2)

		//任务结束
		fmt.Printf("Routine %d completed work %s\n",routineNum,task)
	}
}