package main

import "fmt"

func accumulator() func() int {
	sum := 0
	return func() int {
		sum++
		return sum
	}
}

func main() {
	acc := accumulator()
	fmt.Println(acc()) // 输出：1
	fmt.Println(acc()) // 输出：2
	fmt.Println(acc()) // 输出：3
}
