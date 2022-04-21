package main

import (
	"fmt"
	"go_pl/ch4/github"
	"log"
)

func main() {
	terms := []string{"repo:golang/go", "is:open", "json", "decoder"}
	result, err := github.SearchIssues(terms)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
