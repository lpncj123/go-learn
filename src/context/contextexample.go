package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, sharedData *int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("协程退出")
			return
		default:
			// 使用共享数据
			fmt.Println("共享数据:", *sharedData)
			// 模拟工作
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 创建一个带有取消功能的 context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 在程序退出时调用 cancel 函数，多次调用不影响

	sharedData := 123
	go worker(ctx, &sharedData)

	// 模拟共享数据的修改
	for i := 0; i < 5; i++ {
		sharedData = i
		time.Sleep(time.Second)
	}

	// 取消协程
	cancel()
	time.Sleep(time.Second)
	fmt.Println("主程序退出")
}
