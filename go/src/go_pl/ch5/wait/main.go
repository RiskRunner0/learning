package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// WaitForServer attempts to contact the server of the URL.
// It tries for one minute using exponential backoff
// It reports an error if all attempts fail.
func WaitForserver(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}

		log.Printf("server not responsind (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wait url \n")
		os.Exit(1)
	}

	url := os.Args[1]

	if err := WaitForserver(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}
