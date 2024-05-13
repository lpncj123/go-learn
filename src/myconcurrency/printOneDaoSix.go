package main

import (
	"fmt"
	"sync"
)

// 打印奇数
func PrintOddNumber(wg *sync.WaitGroup, ch chan int, num int) {
	defer wg.Done()
	for i := 0; i <= num; i++ {
		ch <- i
		if i%2 != 0 {
			fmt.Println("奇数：", i)
		}
	}
}

// 打印偶数
func PrintEvenNumber(wg *sync.WaitGroup, ch chan int, num int) {
	defer wg.Done()
	for i := 1; i <= num; i++ {
		<-ch
		if i%2 == 0 {
			fmt.Println("偶数：", i)
		}
	}
}
func PrintOneDaoSix() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	//打印到偶数或者奇数方法循环结束就调用done，主线程结束，终止打印
	wg.Add(1)
	go PrintOddNumber(&wg, ch, 10)
	go PrintEvenNumber(&wg, ch, 10)
	wg.Wait()
	println("打印完毕")
}

func SyncGroupTest() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("hello")
	}()

	wg.Wait()
	println("打印完毕")
}
