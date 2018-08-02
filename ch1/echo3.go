/**
1.数字是怎么定义的？
	Args[0:]
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], ","))
}
