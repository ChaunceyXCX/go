/*
 * @Date: 2024-01-30 16:27:32
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-30 17:38:25
 * @FilePath: \go_learn\ch4\issues\main.go
 * @Description:
 */
package main

import (
	"fmt"
	"go_learn/ch4/github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("issues: %5d\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
