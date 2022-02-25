package main

import "fmt"

func init()  {
	
}

func main()  {
	getIf()
	getSwitch()
}

func getIf(){
	a := 100
	b := 200
	if a > 300{
		fmt.Println("a不大于300")
	}else if b >= 200{
		fmt.Println("b不大于200")
	}else{
		fmt.Println("最后")
	}

	x := 0
	if c := 100; x <= 0{
		fmt.Println("x小于等于0", c)
	}

	fmt.Println("*不支持三元操作符(三目运算符) \"a > b ? a : b\"。")

	if a > 200{
		fmt.Println("大于200")
	} else if a > 150 {
		fmt.Println("大于150")
	} else {
		if b == 200{

		}
		fmt.Println("...")
	}
}

func getSwitch(){
	var marks int = 90

	switch marks {
		case 90:
			fmt.Println(90)
		case 80:
			fmt.Println(80)
		case 50,60,70 :
			fmt.Println(70)
		default:
			fmt.Println("default")
	}
}
