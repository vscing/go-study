package main

import (
	"fmt"
	"strings"
	"reflect"
)

func main(){
	a := 100
	a1 := false
	a2 := 3.1415
	a3 := "hello"
	a4 := 'v'
	fmt.Printf("%T, %T, %T, %T, %T, \n%s \n", a, a1, a2, a3, a4, reflect.TypeOf(a4))

	testString()
	testByte()
}

func testString(){
	str := "hello word"
	fmt.Printf("字符串长度%d\n", len(str))
	fmt.Printf("字符串拼接:\n1.+号 %s\n", str+" 123")
	res := fmt.Sprintf("2.函数Sprintf %s\n", "ccc")
	fmt.Println(res)
	res1 := strings.Split(str, " ")
	fmt.Println("字符串拆分", res1, res1[0], res1[1], len(res1))
	res = strings.Join(res1, " ")
	fmt.Println("字符串合并", res)
	res2 := strings.Contains(str, "123")
	fmt.Printf("字符串%s,不包含123？%t\n", str, res2)
	fmt.Printf("前缀%t，后缀%t\n", strings.HasPrefix(str, "1"), strings.HasSuffix(str, "d"))
	fmt.Println("字符串位置查找", strings.Index(str, "he"), strings.LastIndex(str, "d"))
}

func testByte()  {
	s := "35.com博客"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%c) %T --", r, r, r)
	}
	fmt.Println()
	s1 := "hello"
	//强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS1 := []rune(s2)
	runeS1[0] = '好'
	fmt.Println(string(runeS1))
}