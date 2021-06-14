package main

import (
	"fmt"
	"os"
)

func main() {
	// 记录字符串输入次数的map
	counts := make(map[string]uint32, 0)
	CountLines(os.Stdin, counts)
	for k, v := range counts {
		fmt.Printf("从控制台输入的字符串%s： 循环了 %d 次 \n", k, v)
	}
}
