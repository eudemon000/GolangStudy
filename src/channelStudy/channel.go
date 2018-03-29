package channelStudy

import (
	"fmt"
	"time"
)

func Count(ch chan int) {
	//fmt.Println("执行了")
	ch <- 1
	fmt.Println("Counting")
	//这个方法存在问题，当ch <- 1后，当前协程会阻塞，等待channel里的数值被读取，当主协程里读取channel的值后阻塞结束，
	// 开始执行fmt.Println，由于IO操作属于耗时操作，所以可能会出现主协程里循环全部结束，并退出程序，然后其他协程的打印
	// 操作并未执行完成，所以控制台可能会出现Counting打印次数少于10次
}

func ChannelTest() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		<- ch
	}
	time.Sleep(time.Second * 3)

}

func SelectTest() {
	fmt.Println("这是Select操作")
	chs := make([]chan int, 5)
	for i:= 0; i < 5; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for i := 0; i < 5; i++ {
	select {
	case <- chs[0]:
	case <- chs[1]:
	case <- chs[2]:
	case <- chs[3]:
	case <- chs[4]:
	}
	}
	time.Sleep(time.Second * 3)
}

func SelectTest2() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
			case ch <- 1:
		}
		i := <- ch
		fmt.Println("Value received:", i)
	}
}

//golang本身并不提供超时操作，可以利用select来实现超时操作，首先可以定义一个超时channel，
//再创建一个匿名函数，在函数内通过time.Sleep来设置超时时间。当程序执行过程中ch阻塞了（没有
//读取数据或者写入数据），select里可以通过timeout channel来继续让程序往下执行
func TimeOut() {
	ch := make(chan int)
	//ch <- 1
	timeOut := make(chan bool, 1)
	go func () {
		time.Sleep(time.Second * 15)
		timeOut <- true
	}()

	select {
	case <- ch:
		case <- timeOut:
			fmt.Println("timeout")
	}

}

type PipeData struct {
	Value	int
	Handler func(int) int
	Next	chan int
}

func Handler(queue chan * PipeData) {
	for data := range queue {
		data.Next <- data.Handler(data.Value)
	}
}
