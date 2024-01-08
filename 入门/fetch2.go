/*
 * @Date: 2024-01-08 15:12:44
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-08 15:30:20
 * @FilePath: \go_learn\入门\fetch2.go
 * @Description:为输入链接补全http
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg,"http") && !strings.HasPrefix(arg,"HTTP") {
			arg = "https://"+arg
		}
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// body, err := io.ReadAll(resp.Body)
		fmt.Printf("状态码:%s\n",resp.Status)
		io.Copy(os.Stdout,resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", arg, err)
			os.Exit(1)
		}

	}
}