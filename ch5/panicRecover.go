/*
 * @Date: 2024-02-02 15:03:27
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-02-02 15:07:51
 * @FilePath: \go_learn\ch5\panicRecover.go
 * @Description:
 */
package main

import "fmt"

func panicRecover() (result string) {
	defer func(){
		if p := recover(); p != nil {
			result = p.(string) //断言string
		}
	}()
	panic("a problem")
}

func main() {
	result := panicRecover()
	fmt.Println(result)
}