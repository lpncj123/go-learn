package main

import "fmt"

func main() {
	a := make(map[string]*Customer, 10)
	a["1"] = NewCustomerZz(1, "张三", "", 123, "123@qq.com", "")
	a["2"] = NewCustomerZz(2, "李四", "", 123, "123@qq.com", "")
	//给key为1的value的email赋值
	a["1"].Name = "张三1"

	delete(a, "1")
	//遍历
	for _, customer := range a {
		fmt.Println(customer)
	}
}
