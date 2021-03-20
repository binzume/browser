package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/binzume/browser"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %v get https://www.google.com/", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}

	browser, err := browser.NewWebBrowser()
	if (err) != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "get":
		req, _ := http.NewRequest("GET", os.Args[2], nil)
		resp, _ := browser.Do(req)
		defer resp.Body.Close()
	default:
		printUsage()
	}
}
