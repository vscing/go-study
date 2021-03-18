package main

import (
	"fmt"
	"math/rand"
	"time"
)

//全局数组
var arr [5]int = [5]int{1, 2, 3}
var arr1 = [...]int{1, 2, 3}
var arr2 = [6]int{1, 2, 3}
var arr3 = [5]string{3:"hello", 4:"word"}

func main()  {
	fmt.Println("golang")
	oneArr()
	fmt.Println(arr)
	twoArr()
	a := [2]int{8, 9}
	fmt.Printf("a: %p\n", &a)
	test(&a)
	fmt.Println(a, "求和: ", sumNum(&a))

	arr := [5]int{3, 5, 7}
	sli := arr[1:4]
	//len(sli)表示可见元素有几个  cap(sli)表示所有元素有几个
	fmt.Println(len(sli), cap(sli), len(sli[:cap(sli)]))

	//随机数
	fmt.Println("时间:", time.Now().Unix())
	rand.Seed(time.Now().Unix())
	var b [10]int
	for i := 0; i < len(b); i++ {
		// 产生一个0到1000随机数
		b[i] = rand.Intn(1000)
	}
	for k, v := range b {
		fmt.Printf("(%d)=%d\n", k, v)
	}

	//题目
	b1 := [5]int{1, 3, 5, 8, 7}
	for k, v := range b1{
		for j, v1 := range b1{
			if v+v1 == 8{
				fmt.Printf("元素和为8的下标(%d, %d)\n", k, j)
			}
		}
	}
	myTest(&b1, 8)

}

func oneArr()  {
	//局部
	a := [5]int{1, 2, 3, 4}
	arr := [3]string{0: "sss"}
	arr1 := [...]int{1, 2, 3}
	d := [...]struct{
		name string
		age uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}
	fmt.Println(arr, arr1, d, len(a))
}

var a [5][3]int

func twoArr()  {
	b := [5][3]int{{1,2}, {3,4}}
	c := [...][2]int{{4,5}, {6,7}} //第 2 纬度不能用 "..."
	fmt.Println(a, b, c)

	//遍历
	var arr2 = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	for k, v := range arr2{
		for k1, v1 := range v{
			fmt.Printf("(%d,%d)=%d\n", k, k1, v1)
		}
	}
}

func test(x *[2]int) {
	fmt.Printf("x: %p\n", x)
	*x = [2]int{1, 2}
	fmt.Println(*x)
}

func sumNum(a *[2]int) int  {
	total := 0
	for _, v := range *a {
		total += v
	}
	return total
}

func myTest(a *[5]int, target int) {
	// 遍历数组
	for i := 0; i < len(*a); i++ {
		other := target - a[i]
		// 继续遍历
		for j := i + 1; j < len(*a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
