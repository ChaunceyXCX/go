/*
 * @Date: 2024-01-10 14:21:43
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-10 14:24:54
 * @FilePath: \go_learn\程序结构\tempconv.go
 * @Description: 包学习
 */
package tempconv

import "fmt"

type Celsius float64 //摄氏温度
type Fahrenheit float64 //华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 //绝对零度
	FreezingC Celsius = 0 // 冰点
	BoilingC Celsius = 100 // 沸点
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }