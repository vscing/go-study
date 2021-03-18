package main

import "fmt"

func init()  {
	
}

func main()  {
	a := 100
	b := 200
	if a > 300{
		fmt.Println("a不大于300")
	}else if b >= 200{
		fmt.Println("b不大于200")
	}else{
		fmt.Println("最后")
	}

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
