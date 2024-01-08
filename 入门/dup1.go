/*
 * @Date: 2024-01-04 18:05:15
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-08 10:33:49
 * @FilePath: \go_learn\入门\dup.go
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
	// 从标准输入中读取内容
	input := bufio.NewScanner(os.Stdin)
	// 命令行中按 `CTRL+c` 结束输入
	for input.Scan() { 
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n>1 {
			// 格式化输出
			// %d          十进制整数
			// %x, %o, %b  十六进制，八进制，二进制整数。
			// %f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
			// %t          布尔：true或false
			// %c          字符（rune） (Unicode码点)
			// %s          字符串
			// %q          带双引号的字符串"abc"或带单引号的字符'c'
			// %v          变量的自然形式（natural format）
			// %T          变量的类型
			// %%          字面上的百分号标志（无操作数）
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}
