/*
 * @Date: 2024-01-30 16:08:38
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-30 17:26:59
 * @FilePath: \go_learn\ch4\github\issues.go
 * @Description:
 */
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms," "))
	fmt.Printf("URL: %s\n", q)
	resp, err := http.Get(IssueURL+"?q="+q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil,fmt.Errorf("请求出错: %s", resp.Status)
	}

	var result IssueSearchResult
	// 函数的传入传出都是复制变量并不是直接操作变量所以要吧变量的指针传入才能影响到相关变量
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

