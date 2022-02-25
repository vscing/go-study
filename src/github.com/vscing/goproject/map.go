package main

import "fmt"

func init()  {
	
}

func main()  {
	map1()
}

func map1()  {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 92
	scoreMap["王五"] = 93
	scoreMap["里斯"] = 94
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["里斯"])
	fmt.Printf("type of a:%T\n", scoreMap)

	scoreMap1 := map[string]int{
		"张三": 92,
		"王五": 93,
		"里斯": 94,
	}
	fmt.Println(scoreMap1)
	fmt.Println(scoreMap1["里斯"])
	fmt.Printf("type of a:%T\n", scoreMap1)

	//判断键值是否存在
	val, key := scoreMap["李四"]
	if key{
		fmt.Println("存在", val)
	}else{
		fmt.Println("不存在")
	}

	//循环遍历
	for k, v := range scoreMap{
		fmt.Printf("%s=>%d\n", k, v)
	}

	for k := range scoreMap{
		fmt.Printf("%s=>\n", k)
	}

	for _, v := range scoreMap{
		fmt.Printf("=>%d\n", v)
	}

	//删除键值对
	delete(scoreMap, "里斯")
	fmt.Println(scoreMap)

	//元素为切片
	mapSlice := make([]map[string]string, 3)
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	fmt.Println(mapSlice)
	//值为切片
	sliceMap3 := make(map[string][]string, 3)
	value := make([]string, 0, 2)
	value = append(value, "北京", "上海")
	sliceMap3["中国"] = value
	fmt.Println(sliceMap3)





}
