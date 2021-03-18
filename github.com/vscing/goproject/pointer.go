package main

import "fmt"

func init()  {

}

func main(){
	//缺少初始空间
	//var a *int
	//*a = 100
	//new
	var a *int
	a = new(int)
	*a = 100
	b := 100
	c := &b
	fmt.Println(*a, *c, a, c)
	pointer1()
}

func pointer1()  {
	var a int
	fmt.Println(&a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(a)
}