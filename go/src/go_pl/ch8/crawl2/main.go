package main

import (
	"fmt"
	"go_pl/ch5/links"
	"log"
	"os"
)

// tokens is a counting semaphore used to
// enforce a limit to 20 concurrent requests
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicatedLists

	// Add command line arguments to workList.
	go func() { workList <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { workList <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
