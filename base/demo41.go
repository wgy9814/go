package main

import "fmt"

func main() {

	//演示for-range遍历数组
	heroes := [...]string{"吴广源", "出发点", "脚后跟"}
	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i, v)
		fmt.Printf("heroes[%d]=%v\n", i, heroes[i])
	}
	for _, v := range heroes {
		fmt.Printf("元素的值=%v\n", v)
	}
}
