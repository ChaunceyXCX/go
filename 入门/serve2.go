/*
 * @Date: 2024-01-08 16:43:38
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-08 16:49:54
 * @FilePath: \go_learn\入门\serve2.go
 * @Description:
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	// 浏览器在访问`/count` 时会默认访问网站图标`/favicon.ico`就会路由到`/，这里我们把`count++` 放到`\add`就好了
	http.HandleFunc("/add", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q %q\n", r.URL.Path, r.URL.RawQuery)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}