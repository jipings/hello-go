
package main

import "fmt"

func main() {
	// 创建 map
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")

	for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}
	// 删除元素 delete
	delete(countryCapitalMap, "France");
	fmt.Println(countryCapitalMap)
}