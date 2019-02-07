
package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string // 创建集合
	fmt.Println(countryCapitalMap==nil)
	countryCapitalMap = make(map[string]string) // 如果初始化 map， 那么就会创建一个 nil map。nil map 不能用来存放键值对
	fmt.Println(countryCapitalMap==nil)
	// map 插入key - value对，各个国家对应的首都
	countryCapitalMap["France"] = "Paris";
	countryCapitalMap["Italy"] = "罗马";
	countryCapitalMap["Japan"] = "东京";
	countryCapitalMap["India"] = "新德里";
	// 使用键值输出地图值
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	// 查看元素在集合中是否存在
	captial, ok := countryCapitalMap["美国"] 
	if(ok) {
        fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国的首都不存在")
	}
}