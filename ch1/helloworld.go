/**
package main一开始写成了package ch1，
结果报错“go run: cannot run non-main package”
因为go规定所有可执行的命令必须在main包下面。
换成package main后程序就可以输出“Hello,世界”了
另外go不要求包名唯一，但是要求全路径包名是唯一的就可以了
详细请看：https://golang.org/doc/code.html#PackageNames
 */
package main

import "fmt"

func main() {
	fmt.Println("Hello,世界")
}

