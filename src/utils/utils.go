package utils

import "fmt"

func Count(n1 float64, n2 float64, operator byte) float64 {
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
