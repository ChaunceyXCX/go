/*
 * @Date: 2024-07-28 19:43:29
 * @LastEditors: chauncey chaunceyxcx@gmail.com
 * @LastEditTime: 2024-07-28 20:00:55
 * @FilePath: \go_learn\note\note.go
 * @Description:
 */
package note

import "fmt"

// 2.2 变量与常量
func VariableAndConstant() {
	var v1 int
	var v2 int = 2
	var v3 = 3
	v4 := 4
	v1 = 1

	var (
		v5 = 5
		v6 int = 6
		v7 int
	)

	fmt.Printf("v1=%v,v2=%v,v3=%v,v4=%v %v %v %v", v1, v2, v3, v4, v5, v6, v7)

	const (
		c1 = 8
		c2 = iota
	)
	
}
