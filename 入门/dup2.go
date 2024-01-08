/*
 * @Date: 2024-01-04 18:05:15
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-08 14:24:02
 * @FilePath: \go_learn\入门\dup2.go
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 构造一个<string,int>的键值对,使用hash实现
	// 键可以是任意类型,只要可以用==做比较, 值也可以是任意类型
	counts := make(map[string]int)
	// 读取启动参数
	files := os.Args[1:]
	// 如果启动参数为空不包含要打开的文件, 从标准输入中读取
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup1: %v\n",err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n>1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}


func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Println(f.Name())
		}
	}
}