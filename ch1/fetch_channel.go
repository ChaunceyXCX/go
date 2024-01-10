/*
 * @Date: 2024-01-08 15:32:50
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-09 10:36:28
 * @FilePath: \go_learn\入门\fetch_channel.go
 * @Description:
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	// 创建一个字符串通道用来接收线程返回数据
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// go function : 开启一个goroutine
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	b, err := io.Copy(io.Discard,resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs Read %7d bytes from %s",time.Since(start).Seconds(), b, url)
}
