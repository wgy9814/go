package main

import "fmt"

func main() {

	//string底层是一个byte数组，因此也可以进行切片处理
	str := "hello@atguigu"
	//使用切片获取到atguigu
	slice := str[6:]
	fmt.Println("slice=", slice)

	//string是不可变的，也就是说不能通过 str[0] = 'z'方式来修改字符串
	//str[0] = 'z' [编译不会通过，报错]

	//如果需要修改字符串，可先将string -> []byte / 或者 []rune -> 修改 -》重写转成string
	//hello@atguigu 改成 zello@atguigu
	//arr1 := []byte(str)
	//arr1[0] = 'z'
	//str = string(arr1)
	//fmt.Println("str=", str)

	//细节，我们转成[]byte后，可以处理英文和数字，但是不能处理中文
	//原因是[]byte 字节来处理，而一个汉字是3个字节，会出现乱码
	//解决方法是将 string 转成 []rune 即可 ，因为rune是按字符处理，兼容汉字

	arr1 := []rune(str)
	arr1[0] = '被'
	str = string(arr1)
	fmt.Println("str=", str)
}
