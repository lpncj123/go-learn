package main

/*
指针基本介绍
基本数据类型，变量存的就是值，也叫值类型
获取变量的地址，用&，比如：var num int，获取num的地址&num
指针类型，指针变量存的是一个地址，这个地址指向的空间才存的是值
获取指针类型所指向的值，使用*，比如 *ptr
*/
func main() {
	var num int = 10
	println(&num)
	var ptr *int = &num
	println(*ptr)
	/*
		指针的使用细节
		值类型，都有对应的指针类型，形式为*数据类型
		值类型包括基本数据类型，数组，结构体
		引用类型包括：指针、slice切片、map、管道chan、interface
	*/
	a := 10
	b := 20
	a, b = b, a
	println(a, "****", b)

}
