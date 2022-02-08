package main

import "fmt"

func main() {
	//使用for-range遍历map
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "深圳"
	cities["no3"] = "广州"
	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}
	fmt.Println("cities 有", len(cities), "对key-value")

	//使用for-range遍历一个结构比较复杂的map

	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "赤沙南约新街"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "荷兰"
	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
	}
}
