package main

/*
	包使用的细节：
	1、再给一个文件打包时，该包对应一个文件夹，比如这里的utils文件夹对应的包名就是utils，文件的包名通常和文件所在的文件名一致，一般为小写字母
	2、当一个文件要使用其他函数或变量时，需要先引入对应的包
	3、为了让其他包的文件可以访问到本包的函数，函数名需要首字母大写，类似于其他语言的public，这样才能跨包访问
	4、在访问其他包函数或者变量时，使用包名.函数或变量名
	5、如果包名较长，可以取别名
	6、在同一包下，不能有相同的函数名和相同的全局变量名
	7、如果要编译成一个可执行程序文件，就需要将这个包声明为main，及package main这个就是语法规范

*/
import (
	"fmt"
	"gohelloworld/src/utils"
)

func main() {
	fmt.Println(utils.Count(1.44, 2.11, '+'))
}
