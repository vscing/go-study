package main

import "fmt"

func init()  {
	
}

func main()  {
	//声明切片
	var arr []int
	fmt.Printf("切片元素个数%d, 切片容量%d\n", len(arr), cap(arr))
	if arr == nil {
		fmt.Println("切片是空")
	} else {
		fmt.Println("切片不是空")
	}
	//切片
	s2 := []int{}
	//make
	s3 := make([]int, 0)
	//make 初始化赋值
	s4 := make([]int, 0, 0)
	s5 := []int{1, 2, 3}
	fmt.Println(s2, s3, s4, s5)

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[:]
	for k, v := range slice{
		fmt.Printf("(%d)=%d(%T)\n", k, v, slice)
	}
	var a = []int{1, 2, 3}
	var b = []int{4, 2, 3}
	c := append(a, b...)
	fmt.Println(c)
	slice1 := data[:2:3]
	slice1 = append(slice1, 100, 200) // 一次 append 两个值，超出 s.cap 限制。
	fmt.Println(slice1, data)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&slice1[0], &data[0]) // 比对底层数组起始指针。

	test1()

	test2()

	test5()


}

func test1()  {
	s := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}

func test2()  {
	arrayA := [2]int{100, 200}

	var arrayB [2]int

	arrayB = arrayA

	fmt.Printf("arrayA的内存地址：%p\narrayB的内存地址：%p\n", &arrayA, &arrayB)

	test4(&arrayA)
}

func test3(arrayC [2]int)  {
	fmt.Printf("arrayC的内存地址：%p\n", &arrayC)
}

func test4(arrayC *[2]int)  {
	fmt.Printf("arrayC的内存地址：%p\n", arrayC)
}

func test5(){
	arrayA := [2]int{100, 200}
	//testArrayPoint(&arrayA)
	arrayB := arrayA[:]
	testArrayPoint(&arrayB)
	fmt.Printf("arrayA的内存地址：%p(%v)\narrayB的内存地址：%p(%v)\n", &arrayA, arrayA, &arrayB, arrayB)

}

func testArrayPoint(x *[]int) {
	fmt.Printf("x的内存地址：%p(%v)\n", x, *x)
	(*x)[1] += 100
}

