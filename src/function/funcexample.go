package main

import "fmt"

func main() {
	res := count(1.5, 2, '/')
	fmt.Println("result = ", res)
}

/*
方法的定义

	func 方法名（参数；和java不同，变量类型写在后面） 返回值类型：如果不返回回，为空{
		方法逻辑
		如果有返回值 return 返回值，无返回值不需要
	}
*/
func sum() {
	println(1)
}
func count(n1 float64, n2 float64, operator byte) float64 {
	var result float64
	switch operator {
	case '+':
		result = n1 + n2
	case '-':
		result = n1 - n2
	case '*':
		result = n1 * n2
	case '/':
		result = n1 / n2
	default:
		fmt.Println("操作符不存在")
	}
	return result
}
