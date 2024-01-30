/*
 * @Date: 2024-01-30 16:09:14
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-30 18:05:48
 * @FilePath: \go_learn\ch4\github\github.go
 * @Description:
 */
package github

import "time"

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int64 `json:"total_count"`
	Items []*Items
}

type Items struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"create_at"`
	Body string
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}