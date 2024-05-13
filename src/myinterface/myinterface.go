package main

type Animal interface {
	Sleep()
	Eat(food string)
}

// Dog 定义dog实现Animal接口，重写Sleep和Eat方法
type Dog struct {
	Name string
	Age  int
}

// Sleep 不要求类型显示地声明实现了某个接口，只要实现了相关的方法即可，编译器就能检测到。
func (d Dog) Sleep() {
	println("Dog Sleep")
}
func (d Dog) Eat(food string) {
	println("Dog Eat", food)
}

//func main() {
//	var a Animal
//	a = Dog{"Dog", 12}
//	a.Sleep()
//	a.Eat("Bone")
//}
