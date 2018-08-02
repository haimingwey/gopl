/**
1.不能引入多余的包，不然会编译不通过
2.GoLand会自动添加所需要的包，代码中直接应用即可
3.声明变量没有显式初始化时，会隐式赋零值，数值型的为0，字符串型为""空字符串
4.发现len 函数是内建函数（built-in function），那么内建函数的源码怎么看呢？
其实这个问题本身有问题，你以为这些内建函数都是有源码的，但是这些是go编译器内建的函数，
所以，我们需要研究编译器是如何实现这些内建函数的，是怎么处理的。
5.go是否有运行时？ TODO
	有一个运行时库，但是和java一般意义上的运行时环境（虚拟机）不一样，go通过编译，
直接变成本地的机器码（参考：https://golang.org/doc/faq#runtime）
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	var cnt = len(os.Args)
	for i := 1; i < cnt; i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
