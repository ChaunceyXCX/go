/*
 * @Date: 2024-01-30 16:55:45
 * @LastEditors: chaunceyxie chaunceyxcx@gmail.com
 * @LastEditTime: 2024-01-31 17:06:55
 * @FilePath: \go_learn\ch4\issuetext\texttemp.go
 * @Description:
 */
package main

import (
	"go_learn/ch4/github"
	"text/template"
	"log"
	"os"
	"time"
)

var report = template.Must(template.New("issueList").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(`
	{{.TotalCount}} issues:
	{{range .Items}}----------------------------------------
	Number: {{.Number}}
	User:   {{.User.Login}}
	Title:  {{.Title | printf "%.64s"}}
	Age:    {{.CreatedAt | daysAgo}} days
	{{end}}
	`))

func main() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(date time.Time) float64 {
	return time.Since(date).Hours() / 24
}