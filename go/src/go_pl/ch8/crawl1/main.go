package main

import (
	"fmt"
	"go_pl/ch5/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)

	// Start with the command-line arguments
	go func() { workList <- os.Args[1:] }()

	// Crawl the web concurrently
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}
