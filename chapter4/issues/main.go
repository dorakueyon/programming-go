package main

import (
	"fmt"
	"github.com/dorakueyon/programming-go/chapter4/github"
	"log"
	"os"
	"time"
)

var withinOneMonthIssues []*github.Issue
var withinOneYearIssues []*github.Issue
var overOneYearIssues []*github.Issue

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if item.CreatedAt.Unix() < time.Now().AddDate(-1, 0, 0).Unix() {
			overOneYearIssues = append(overOneYearIssues, item)
		} else if item.CreatedAt.Unix() < time.Now().AddDate(0, -1, 0).Unix() {
			withinOneYearIssues = append(withinOneYearIssues, item)
		} else {
			withinOneMonthIssues = append(withinOneMonthIssues, item)
		}

	}
	fmt.Println("within this month")
	for _, item := range withinOneMonthIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("within a year")
	for _, item := range withinOneYearIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("long long time ago...")
	for _, item := range overOneYearIssues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
