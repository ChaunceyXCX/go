/*
 * @Date: 2024-02-01 15:03:56
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-02-01 15:08:30
 * @FilePath: \go_learn\ch5\sum.go
 * @Description:
 */
package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func maxMin(vals ...int) (int, int, error) {
	max, min := 0, 0
	if len(vals) == 0 {
		return max, min, fmt.Errorf("no values")
	}
	for _, val := range vals {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return max, min, nil
}

func main() {
	fmt.Println(sum())
	fmt.Println(maxMin(6,5,4,7,8,9))
}