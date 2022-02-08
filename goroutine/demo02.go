package main
import(
	"fmt"
	"runtime"
)

func main(){
	//获取当前系统逻辑cpu的数量
	num := runtime.NumCPU()
	//GOMAXPROCS设置可同时执行的最大CPU数，并返回先前的设置
	runtime.GOMAXPROCS(num)
	fmt.Println("num=", num)
}
//go1.8后，默认让程序运行在多个核上，可以不用设置了
//go1.8前，还是要设置一下，可以更高效的利益cpu