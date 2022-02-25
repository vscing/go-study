package main

import "fmt"

var _ int64=i()

var name = 100 //全局的

const pi = 3.1415 //常量
//批量声明
const (
	e = 2.7777
	b
	c
)

//批量声明
var (
	a1 string
	a2 int
	a3 bool
	a4 float32
)

func init(){
	fmt.Println("init函数")
	fmt.Println("常量:", pi)
	fmt.Println("批量常量 如果省略了值则表示和上面一行的值相同", e, b, c)
}

const (
	n1 = iota
	n2
	n3
	n4
)

const (
	n5 = iota
	n6 = 100
	_  = iota
	n7
)
//多个iota定义在一行
const (
	n8, n9 = iota+1, iota+2
	n10, n11
	n12, n13
)

func init(){
	name := 100
	fmt.Println("init1函数", name)
	fmt.Println("常量计数器 iota", n1, n2, n3, n4)
	fmt.Println("常量计数器间隔", n5, n6, n7)
	fmt.Println("常量计数器多个定义", n8, n9, n10, n11, n12, n13)
}

func main(){
	var a = 1
	var b = "hello word" //类型推导
	var name, age = "张三", 10 //初始化多个变量
	sex := "男" //短变量声明 函数内使用

	s()
	x, _ := foo() //匿名变量
	_, y := foo() //匿名变量


	fmt.Println("str := \"c:\\pprof\\main.exe\"")
	fmt.Println(`第一行
		多行字符串
		123\n
		3456\"
	最后一行`)
	fmt.Println(a, b, name, age)
	fmt.Println("匿名变量:", x, y)
	fmt.Printf("%s今年大概是%d,性别是%s", name, age, sex)
}

func i() int64{
	fmt.Println("s函数", name)
	return 1
}

func s() string{
	name := 100
	fmt.Printf("返回一个字符串: %d \n", name)
	return ""
}

func foo() (int, string) {
	return 10, "曹"
}