package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//1.统计字符串的长度，按字节 len(str)
	//golang的编码统一为utf-8(ascii的字符（字母和数字）占一个字节，汉子占3个字节)
	str := "hello北"
	fmt.Println("str len=", len(str))

	str2 := "hello北京"
	//2.遍历字符串，同时处理有中文的问题 r := []rune(str2)
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n",r[i])
	}

	//3.字符串转整数： n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果是", n)
	}

	//4.整数转字符串 str := strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T\n", str, str)

	//5.字符串转 []byte: var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("str=%v \n", bytes)

	//6.[]byte 转字符串 : str = string([]byte{97,98,99})
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v \n", str)

	//7. 10进制转2 8 16进制： str = strconv.FormatInt(123, 2),返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制是=%v\n", str)

	//8.查找子串是否在指定的字符串中： strings.Contains("seafood","foo") true
	b := strings.Contains("seafood","foo")
	fmt.Printf("b=%v \n", b)

	//9.统计一个字符串中有几个指定的子串 ：strings.Count("ceheese", "e") //4
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v \n", num)

	//10.不区分大小写的字符串比较（==是区分大小写的）: strings.EqualFold("abc", "Abc")
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v \n", b)
	fmt.Println("结果", "abc" == "Abc") //false 区分大小写

	//11.返回子串在字符串第一次出现的index值，如果没有返回-1
	//strings.Index("NLT_abc", "abc") // 4
	index := strings.Index("NLT_abc", "abc")
	fmt.Printf("index=%v \n", index)

	//12.返回子串在字符串最后一次出现的index，如果没有返回-1
	//strings.LastIndex("go golang", "go")
	index = strings.LastIndex("go golang", "go") //3
	fmt.Printf("index=%v \n", index)

	//13.将指定的字符串换成 另一个子串 ： strings.Replace("go go hello", "go", "广州", n)
	//n可以指定你希望替换几个，如果n=-1表示全部替换
	str2 = "go go hello"
	str = strings.Replace(str2, "go", "广州", -1)
	fmt.Printf("str=%v str2=%v\n", str, str2)

	//14.按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：
	//strings.Split("hello,world,ok", ",")
	strArr := strings.Split("hello,world,ok", ",")
	for i :=0; i < len(strArr); i++{
		fmt.Printf("str[%v]=%v \n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)

	//15.将字符串的字母进行大小写的转换
	//strings.ToLower("Go") //go strings.ToUpper("Go") // GO
	str = "golang hello"
	str = strings.ToLower(str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str)

	//16.将字符串左右两边的空格去掉：	strings.TrimSpace(" i am a doctor ")
	str = strings.TrimSpace(" i am a doctor ")
	fmt.Printf("str=%q\n", str)

	//17.将字符串左右两边指定的字符去掉：
	//strings.Trim(" !hello!", "!") ["hello"] //将左右两边!和空格去掉
	str = strings.Trim("! hello! ", " !")
	fmt.Printf("str=%q\n", str)

	//20.判断字符串是否以指定的字符串开头
	//strings.HasPrefix("ftp://192.168.11", "ftp") //true
	b = strings.HasPrefix("ftp://192.168.11", "ftp")
	fmt.Printf("b=%v \n", b)

	//21.判断字符串是否以指定的字符串结束
	//strings.HasSuffix("ftp://192.168.11", "ftp") //true
	b = strings.HasSuffix("ftp://192.168.11", "ftp")
	fmt.Printf("b=%v \n", b)
}
