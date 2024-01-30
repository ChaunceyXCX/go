/*
 * @Date: 2024-01-22 16:30:36
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-30 16:07:43
 * @FilePath: \go_learn\ch3\comma_float.go
 * @Description:
 */
package comma

import "strings"

func Comma_float(s string) string {
	// 分别处理小数点前后的
	// 小数点前面的
	before := s[:strings.LastIndex(s, ".")]
	// 小数点后面的
	after := s[strings.LastIndex(s, ".")+1:]
	return Comma(before) + "." + Comma(after)
}

func commaAfter(s string) string{

	return ""
}
