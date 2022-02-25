package main

import (
	"encoding/json"
	"fmt"
)

func init()  {
	
}

//类型定义
type NewInt int

//类型别名
type MyInt = int

//结构体
type person struct{
	name string
	age int
	city string
}

type Person struct {
	name string
	age int
	city string
}

func main()  {
	var a NewInt
	var b MyInt

	a.SayHello()
	fmt.Printf("type of a:%T b:%T\n", a, b)

	var p1 person
	fmt.Printf("结构体初始化=%#v\n", p1)
	p1.age = 30
	p1.city = "北京"
	p1.name = "小明"
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)

	//匿名结构体
	var user struct{Name string; Age int}
	user.Age = 20
	user.Name = "张三"
	fmt.Printf("user=%v\n", user)
	fmt.Printf("user=%#v\n", user)

	//指针类型结构体
	var p2 = new(person)
	p2.age = 30
	p2.city = "北京"
	p2.name = "小明"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", &p2)
	fmt.Printf("%#v\n", p2)
	fmt.Println(p2, *p2)

	//取结构体的地址实例化
	p3 := &person{}
	p3.age = 30
	p3.city = "北京"
	p3.name = "小明"
	fmt.Printf("%T\n", p3)
	fmt.Printf("%p\n", &p3)
	fmt.Printf("%#v\n", p3)
	fmt.Println(p3, *p3)

	//使用键对值初始化结构体
	p4 := person{
		name: "王五",
		age: 20,
		city: "北京",
	}
	fmt.Println(p4)

	//使用键对值初始化指针结构体
	p5 := person{
		name: "王五",
		age: 20,
		city: "北京",
	}
	fmt.Println(p5)
	//初始化部分
	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7)

	//使用值的列表初始化
	p8 := person{
		"王五",
		20,
		"北京",
	}
	fmt.Printf("p8=%#v\n", p8)

	//结构体内存布局
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	p9 := newPerson("pprof.cn", "测试", 90)
	fmt.Printf("%#v\n", p9)

	//方法和接收者
	p11 := newPerson("pprof.cn","测试", 25)
	p11.Dream()
	fmt.Println(p11.age)
	p11.SetAge(100)
	fmt.Println(p11.age)

	p12 := Person2{
		"pprof.cn",
		18,
	}
	fmt.Printf("%#v\n", p12)        //main.Person{string:"pprof.cn", int:18}
	fmt.Println(p12.string, p12.int)

	user1 := User{
		Name:   "pprof",
		Gender: "女",
		Address: Address{
			Province: "黑龙江",
			City:     "哈尔滨",
		},
	}
	fmt.Printf("user1=%#v\n", user1)

	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang()
	d1.move()

	json1()

	var ce []student  //定义一个切片类型的结构体
	ce = []student{
		student{1, "xiaoming", 22},
		student{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)
}

func newPerson(name string, city string, age int) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// 使用指针接收者
func (p *Person) SetAge(newAge int) {
	p.age = newAge
}

func (m NewInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

//结构体的匿名字段
type Person2 struct {
	string
	int
}

//嵌套结构体
type Address struct {
	Province string
	City     string
}
//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

//结构体的“继承”
//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

//结构体字段的可见性
//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

//Student 学生
type Student struct {
	ID     int `json:"id"`
	Gender string
	name   string //私有不能被json包访问
}
//Class 班级
type Class struct {
	Title    string
	Students []*Student
}
//json序列化
func json1()  {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			//Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil{
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

//练习
type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	// ce = append(ce, student{3, "xiaowang", 56})
	// return ce
}