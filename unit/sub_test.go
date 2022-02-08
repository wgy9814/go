package main

import (
	"fmt"
	"testing"
)

//确保每个函数是可运行，并且运行结果是正确的
//确保写出来的代码性能是好的，
//单元测试能及时的发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决，而性能测试的重点在于发现程序设计上的一些问题，让程序能够在高并发的情况下还能保持稳定

//总结
//测试用例文件名必须以 _test.go 结尾。 比如 cal_test.go 。
//测试用例函数必须以 Test 开头，一般来说就是 Test+被测试的函数名，比如 TestAddUpper
//TestAddUpper(t *tesing.T)的形参类型必须是 *testing.T
//一个测试用例文件中，可以有多个测试用例函数，比如 TestAddUpper、TestSub

func TestGetSub(t *testing.T) {

	//调用
	res := getSub(50,20)
	if res != 30 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
		t.Fatalf("GetSub(50,20) 执行错误，期望值=%v 实际值=%v\n", 30, res)
	}
	//如果正确，输出日志
	t.Logf("GetSub(50,20) 执行正确...")
}

func TestHello(t *testing.T) {
	fmt.Println("TestHello正在被调用")
}