/*
 * @Date: 2024-01-09 14:54:50
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-09 15:28:31
 * @FilePath: \go_learn\入门\echo4.go
 * @Description:
 */
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "禁止打印换行符")
var s = flag.String("s", " ", "自定义字符串分割符")

func main() {
	// 解析命令行参数
	flag.Parse()
	// s是一个指针类型所以要加型号取得变量值
	fmt.Print(strings.Join(flag.Args(), *s))
	if !*n {
		fmt.Println() // 打印换行符 powershell 可能会吞掉换cmd才能看到效果
	}
}