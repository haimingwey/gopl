/**
输入并且打印出来
*/
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	//新建一个map，用来存储输入
	count := make(map[string]int)
	//定义输入流
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "q" {
			break
		}
		count[input.Text()]++
	}
	//读取输入流
	for line, n := range count {
		fmt.Printf("%d\t%s\n", n, line)
	}
	//打印输出
}
