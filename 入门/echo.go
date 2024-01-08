/*
 * @Date: 2024-01-03 17:52:42
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-04 18:04:21
 * @FilePath: \go-hello-world\echo\echo.go
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[0:] {
		if index == 0 {
			fmt.Println("os.Args[0] =", arg)
		} else {
			fmt.Println("index:", index, ", value:", arg)
		}
	}
}
