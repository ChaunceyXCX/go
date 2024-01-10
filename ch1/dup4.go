/*
 * @Date: 2024-01-08 14:23:28
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-08 14:37:51
 * @FilePath: \go_learn\入门\dup4.go
 * @Description: 练习1.4 打印重复行对应的文件名称
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
	countsWithFileName := make(map[string]map[string]int)
	// 读取启动参数
	files := os.Args[1:]
	// 如果启动参数为空不包含要打开的文件, 从标准输入中读取
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines(os.Stdin, counts, nil)
	} else {
		for _, arg := range files {
			// 每个文件有自己独立的counts map
			counts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}
			countLines(f, counts, countsWithFileName)
			f.Close()
		}
	}
	for fileName, counts := range countsWithFileName {
		for line, count := range counts {
			if count > 1 {
				fmt.Printf("%s\t%d\t%s\n", fileName, count, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, countsWithFileName map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	if countsWithFileName != nil {
		countsWithFileName[f.Name()] = counts
	}
}
