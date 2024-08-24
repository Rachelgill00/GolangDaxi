package main

import "fmt"

// Producer2返回的是一个"只接受通道"
// 这就从代码层面限制了该函数返回的通道只能进行“接受操作”， 保证了数据安全
func Producer2() <-chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行 发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后 关闭通道
	}()

	return ch
}

// Consumer从通道中接受数据进行计算
func Consumer2(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}

	return sum
}

func main() {
	ch := Producer2()

	res := Consumer2(ch)
	fmt.Println(res)

	var ch3 = make(chan int, 1)
	ch3 <- 10
	close(ch3)
	Consumer2(ch3) // 函数传参时将ch3“转为单向通道”

	var ch4 = make(chan int, 1)
	ch4 <- 10
	/*
		var ch5 <-chan int // 声明一个“只接收通道”
		ch5 = ch4          // 变量赋值时将ch4转为单向通道
	*/
	ch5 := (<-chan int)(ch4)
	<-ch5
}
