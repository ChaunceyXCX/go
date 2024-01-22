/*
 * @Date: 2024-01-22 16:41:53
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-22 17:00:36
 * @FilePath: \go_learn\ch3\comma_buf.go
 * @Description:
 */
package comma

import (
	"bytes"
	"strings"
)

func CommaBuf(s string) string {
	// 小数点前面的
	before := s[:strings.LastIndex(s, ".")]
	// 小数点后面的
	after := s[strings.LastIndex(s, ".")+1:]

	return commaBufCommon(before) + "." + commaBufCommon(after)
}

func commaBufCommon(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
		
	}
	return buf.String()
}