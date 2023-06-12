package main

import (
	"fmt"
	"time"
)

/*
	Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
	goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
	语法格式：go 函数名( 参数列表 )
*/
//主程序不会等待线程执行结束就退出了
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/*
通道(channel)
通道（channel）是用来传递数据的一个数据结构。
通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。

ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据  ch := make(chan int)

	// 并把值赋给 v
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

/*通道关闭*/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	/*	go say("hello gosay")
		say("hello")*/

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	fmt.Printf("%v", s[:len(s)/2])
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x := <-c
	y := <-c
	println(x, ",", y, ",", x+y)

	//	通道设置缓冲区
	ch := make(chan int, 2)

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	k := make(chan int, 10)
	go fibonacci(cap(k), k)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range k {
		fmt.Println(i)
	}
}
