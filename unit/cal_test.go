package main

import (
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

//1，测试单个文件，一定要带上被测试的原文件
//    go test -v  cal_test.go cal.go
//
//2，测试单个方法
//
//    定位到被测试的文件路径
//    go test -v -test.run TestAddUpper

func TestAddUpper(t *testing.T) {

	//调用
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
	}
	//如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确...")
}

