package main

import "fmt"

func main() {

	//使用常规的for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	//slice := arr[1:4]//20 30 40
	slice := arr[:] // 等价于 arr[0:len(arr)]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v ", i, slice[i])
	}

	fmt.Println()
	//使用for--range方式遍历切片
	for i, v := range slice {
		fmt.Printf("i=%v v=%v \n", i, v)
	}

	slice2 := slice[1:2]
	slice2[0] = 100 //因为arr slice slice2指向的数据空间都是同一个，改变其他都会改变
	fmt.Println("arr=", arr)
	fmt.Println("slice=", slice)
	fmt.Println("slice2=", slice2)

	fmt.Println()

	//append本质就是对数组扩容
	//go底层会创建一个新数组 newarr（安装后扩容大小）
	//将slice原来包含的元素拷贝到新的数组newarr
	//slice重新引用到newarr

	//用append内置函数，可以对切片进行动态追加
	var slice3 []int = []int{100,200,300}
	//通过append直接给slice3追加具体的元素
	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("slice3=", slice3)

	//通过append将切片slice3追加给slice3
	slice3 = append(slice3, slice3...)
	fmt.Println("slice3=", slice3)

	//切片的拷贝操作
	//copy（1,,2）参数的数据类型是切片
	//切片使用copy内置函数完成拷贝，举例说明
	//slice4和slice5的数据空间是独立的，不互相影响
	fmt.Println()
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)

	fmt.Println("slice4=", slice4)
	fmt.Println("slice5=", slice5) //[1 2 3 4 5 0 0 0 0 0]

	var a []int = []int{1, 2, 3, 4, 5}
	var b = make([]int, 1)
	copy(b, a)

	fmt.Println("a=", a)
	fmt.Println("b=", b) //[1 2 3 4 5 0 0 0 0 0]

}
