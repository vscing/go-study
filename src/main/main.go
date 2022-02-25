package main // 声明 main 包，表明当前是一个可执行程序

import "fmt" // 导入内置 fmt

// init可以有多个，同一文件里执行顺序不确定，其它包文件的init执行顺序按包引入的顺序。
// main只有一个
// go build 编译得到的可执行文件会保存在执行编译命令的当前目录下
// go build -o 名字.exe 还可以使用-o参数来指定编译后可执行文件的名字

func init() {
	fmt.Println("main init1")
}

func init() {
	fmt.Println("main init2")
}

func main() {
	fmt.Println("hello word")
}
