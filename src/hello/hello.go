package hello

import "fmt"

/**
test
*/
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

// 全局变量
var m = 100

func main1111() {
	fmt.Println("hello word", m)
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	/**
	* Go 中 main 包默认不会加载其他文件， 而其他包都是默认加载的。
	* 如果 main 包有多个文件，则在执行的时候需要将其它文件都带上，
	* 即执行 go run .
	 */
	// aa()
}
