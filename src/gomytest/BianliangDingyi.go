package main

import (
	"fmt"
	"gohelloworld/src/gomytest/utils"
)

func main() {
	var a int
	var b int
	a, b = 1, 2
	utils.Call(1, 2)
	call(a, b)
	fmt.Println(a + b)
}
func call(a int, b int) {
	fmt.Println(a, "哈哈", b)
}
