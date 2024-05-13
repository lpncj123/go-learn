// package main
//
// import "fmt"
//
//	func main() {
//		ch := make(chan int, 2)
//
//		ch <- 1
//		x := <-ch
//		ch <- 2
//		y := <-ch
//		// 获取这两个数据
//		ch <- 3
//		fmt.Println(x, y)
//
//		//PrintOneDaoSix()
//	}
package main

import "fmt"

func main() {
	s := []int{1, 1, 1}
	f(s)
	fmt.Println(s)
}

func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	/*for _, i := range s {
		i++
	}
	*/

	for i := range s {
		s[i] += 1
	}
}
