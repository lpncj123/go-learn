package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

/*
定义全局变量
在go中函数外面就是全局变量
*/
var b = 12
var c = "我们"
var (
	d = "666"
	e = "777"
)

type person struct {
	name string
	age  byte
	sex  byte
}

func main() {
	//定义变量 声明类型 一共三种
	//第一种，其他的是在声明时同时赋值就可以省略类型和var
	//var a int
	/*第二种，自动推断类型*/
	//var a = 12.11
	/*第三种，省略var，根据等号后面自动推导类型,注意不能随便改变数据类型，在给变量赋值时候就固定了类型*/
	a := 23
	/*第四种 一次声明多个变量*/
	name, age, sex := "lp", 12, "男"
	fmt.Println(name, age, sex)
	//赋值
	a = 12
	fmt.Println("a=", a)

	/*
		整型的使用细节，golang各整数类型分：有符号和无符号,int uint的大小和系统有关
		整型的默认申明为int型
		golang程序中在变量整型变量在使用时，遵循保小不保大原则，节省更多的内存空间
	*/
	/*查看某个变量的数据类型及占用空间大小*/
	var n int64 = 10
	fmt.Printf("n2的类型为 %T，占用的字节数为 %d", n, unsafe.Sizeof(n))

	/*比如定义年龄*/
	var age1 byte = 23
	fmt.Println(age1)

	/*
		小数类型，浮点型
		分类float32和float64，和java中的float和double
	*/
	num1 := 5.12
	num2 := .123 //0.123
	fmt.Println("num1=", num1, "\n", num2)

	/*
		字符类型，go中没有专门的字符类型，要存储单个字符(字母)，一般使用byte存储，字符串就是一串固定长度的字符链接起来的字符序列。go中的字符串是由字节组成的
		字符的本质就是一个整数，直接输出，输出的为对应的utf-8的码值，%c可以转化为对应的数值
		字符类型相当于一个整数，可以进行计算
		go语言的编码统一都改为了utf-8

	*/
	var c1 byte = 'a'
	var c2 byte = 'o'
	fmt.Println(c1, ",", c2)
	fmt.Printf("c1= %c c2= %c", c1, c2)

	//var c3 byte = '北' //overflows溢出
	var c3 int64 = '北'
	println()
	fmt.Printf("c3= %c c3对应码值为 %d", c3, c3)
	println()

	/*bool类型*/
	var b1 = false
	fmt.Println(unsafe.Sizeof(b1))

	/*string类型 注意的是在java中string是引用数据类型，而这里是基本数据类型
	go语言中的字符串的字节使用了utf-8的编码标识unicode文本
	字符串赋值后了不能再改变，go中字符串是不可变得，就是对字符串操作了，会生成一个新的字符串
	字符串的两种表现形式
	*/
	var address = "上海闵行区\nfd"
	fmt.Println(address)

	/*
		字符串的两种表现形式
		双引号，识别转义字符
		反引号，原样输出字符串
	*/
	var address1 = `\n我们`
	println(address1)

	var address2 = "我" + "你"
	println(address2)

	//基本类型默认值  boolean为false，字符串为""，其他都为0

	/*
			类型转换
		GO不能自动转换，只能显式转换
	*/
	var aa int32 = 12
	var aa1 int64
	var aa2 int8
	aa1 = int64(aa) + 20
	aa2 = int8(aa) + 20
	println(aa, aa1, aa2)

	/*
		转换为字符串
	*/
	z := 99
	z1 := 99.99
	z2 := true
	z3 := 'h'
	var str string

	str = fmt.Sprintf("%d", z)
	fmt.Printf("str type %T,str=%q", str, str)
	str = fmt.Sprintf("%f", z1)
	fmt.Printf("str type %T,str=%q", str, str)
	str = fmt.Sprintf("%t", z2)
	fmt.Printf("str type %T,str=%q", str, str)
	str = fmt.Sprintf("%c", z3)
	fmt.Printf("str type %T,str=%q", str, str)

	str = strconv.FormatInt(int64(z), 10)
	fmt.Printf("str type %T,str=%q", str, str)

	i, _ := strconv.ParseInt(str, 10, 64)
	println(i)

	//数组的定义  var variable_name [SIZE] variable_type
	/*var nums [5]int
	nums = [5]int{1, 2, 3, 4}*/
	/*nums:=[5]int{1,2}*/
	/*var nums = [5]int{1,2}*/

	//结构体的使用
	/*p := person{age: 12, name: "lp", sex: 1}*/
	/*	var kid person
		kid.age = 12
		kid.name = "lp"
		kid.sex = 1
		getPerson(kid)
		getPerson1(&kid)*/

	//Go 语言切片(Slice) 定义：var identifier []type  切片不需要定义长度，就类似与java的集合 var slice1 []type = make([]type, len)  slice1 := make([]type, len)
	//切片的容量大小，就是初始的长度减去去除开头元素的个数
	/*s := []int{1, 2, 3}*/
	/*arr := [4]int{1, 2, 3, 4}*/
	//数组转换为切片
	/*s:= arr[:]*/
	/*startIndex := 1
	s := arr[startIndex:] //这两个参数如果只有startIndex就是从这里截取到最后，只有endIndex，就是从一开始截取到这里，如果都有值，那么是从start-》end，不包括startIndex
	for i2 := range s {
		println(s[i2])
	}
	println(len(s))
	println(cap(s))
	//判断切片是否为空
	if s == nil {
		println("切片为空")
	}*/
	//append方法和copy方法,
	/*var numbers []int
	numbers = append(numbers, 1, 2)
	numbers1 := make([]int, len(numbers), cap(numbers)*2)
	numbers1 = append(numbers1, numbers...) //...用于将切片（或数组）展开成单独的元素序列
	copy(numbers1, numbers)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers1), cap(numbers1), numbers1)*/

	//range范围的使用
	/*	s := []int{1, 2, 3}
		for i2 := range s {
			println(s[i2])
		}
		for index, value := range s {
			fmt.Printf("index= %d,value= %d", index, value)
		}
		for _, i2 := range s {
			println(i2)
		}
		for i, c := range "go" {
			println(i, "****", c)
		}*/

	//map集合的使用  定义  map_variable := make(map[KeyType]ValueType, initialCapacity)
	//定义空的map
	/*map1 := make(map[string]int)*/
	//直接定义map
	/*m := map[string]int{
		"apple": 1,
		"banana": 2,
		"orange": 3,
	}*/
	map1 := make(map[string]int, 10)
	map1["Google"] = 1
	map1["baidu"] = 1000
	for key := range map1 {
		println(map1[key])
	}
	for key, value := range map1 {
		println(key, "*****", value)
	}
	//修改键值对
	map1["Google"] = 0
	println(len(map1))
	//删除键值对
	delete(map1, "Google")

}

/*func getPerson(p person) {
	fmt.Println(p.name)
}

func getPerson1(p *person) {
	fmt.Println(p.age)
}
*/
