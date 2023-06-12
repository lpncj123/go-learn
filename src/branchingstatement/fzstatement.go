package main

func main() {
	//if else
	/*age := 19
	if age > 18 {
		println("大于18")
	} else if age < 18 {
		println("小于等于18")
	} else {

	}*/

	//switch语句
	/*	var key byte
		fmt.Println("请输入一个字符！")
		fmt.Scanf("%c", &key)
		switch addOne(key) + 1 {
		case 'a', 'e':
			fmt.Println("first")
		case 'b':
			fmt.Println("second")
		case 'c':
			fmt.Println("third")
		default:
			fmt.Println("没有")

		}*/

	//for循环
	/*	for i := 1; i <= 10; i++ {
			println("我们")
		}
		//等价于for ;;{},死循环
		for {
			println("aaaaaaaaaaaaaa")
		}*/

	//遍历字符串,这种方式，是相当于字符方式遍历。及时字符串有中文也是可以的。
	//str := "abc~我们的世界"
	/*	for i, i2 := range str {
			fmt.Printf("index = %d,val = %c", i, i2)
		}
		println()*/
	//这样传统的的方式按照字节遍历字符串会出现乱码问题，因为在go语言中，字符串的底层编码为ut-8，一个汉字对应的是三个字节，解决方式：需要将str转换为[]rune切片
	/*	for i := range str {
		fmt.Printf("val = %c\n", str[i])
	}*/

	/*str2 := []rune(str)
	for i := range str2 {
		fmt.Printf("val = %c\n", str2[i])
	}*/

	//注意 go语言中没有while和do while循环，可以用for实现
	//var i int = 1
	//while实现
	/*for {
		if i > 10 {
			break
		}
		println(i)
		i++
	}
	println(i)*/

	//do while实现
	/*for {
		println(i)
		i++
		if i > 10 {
			break
		}
	}
	println(i)*/

	//goto语句 可以无条件的转移到程序中所指定的行， 通常与条件语句配合使用，一般不推荐使用，也就是不使用,return是跳出当前方法
	/*var i int = 30
		if i > 20 {
			goto label
		}
		println(1)
		println(2)
	label:
		println(3)
		for i := 0; i < 5; i++ {
			if i == 3 {
				return
			}
			println(i)
		}*/
}
func addOne(char byte) byte {
	return char + 1
}
