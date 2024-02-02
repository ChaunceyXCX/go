/*
 * @Date: 2024-02-02 15:52:41
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-02-02 15:55:32
 * @FilePath: \go_learn\ch6\methodPTR.go
 * @Description:
 */
package main

import "fmt"

type Point struct {
	x, y int
}

func (p *Point) ScaleBy(factor int) {
	fmt.Println("p:", p)
	p.x *= factor
	p.y *= factor
}

func main() {
	p := Point{1, 2}
	p.ScaleBy(2)
	fmt.Println(p)
}