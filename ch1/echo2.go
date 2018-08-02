/**
1.赋值给空标识符“_”会被丢弃
2.符号“:=”是短变量声明
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
