/*
 * @Date: 2024-01-08 10:59:06
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-08 14:17:58
 * @FilePath: \go_learn\入门\dup3.go
 * @Description:
 */
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		// 从1.16版本开始废弃 用os.ReadFile 替代
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		// split返回的是什么为什么最后一行没有换行符没有计算
		// windows换行符是 "\r\n" unix 是 "\n"
		for index, line := range strings.Split(string(data), "\r\n") {
			fmt.Printf("%d\t%v\n",index,line)
			counts[line] ++
		}
	}
	for line, n := range counts {
		if n>1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}