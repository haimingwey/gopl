/**
1.不能引入多余的包，不然会编译不通过
2.GoLand会自动添加所需要的包，代码中直接应用即可
3.声明变量没有显式初始化时，会隐式赋零值，数值型的为0，字符串型为""空字符串
4.发现len 函数是内建函数（built-in function），TODO 那么内建函数的源码怎么看呢？
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
