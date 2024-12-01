package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/aquasecurity/table"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatalln("missing URL when trying to run the app")
	}

	targetURL := args[1]
	if !IsValidURL(targetURL) {
		log.Fatalln("provided link is not valid")
	}

	linksToVisit := []string{}
	PopulateLinks(targetURL, &linksToVisit)

	if len(linksToVisit) <= 0 {
		fmt.Println("could not scrape any links from that URL. Exiting gracefully")
		os.Exit(0)
	}

	fmt.Printf("found %d links on %s\nChecking...\n", len(linksToVisit), targetURL)

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	t := table.New(os.Stdout)
	current := 1

	t.SetHeaders("#", "link", "status")

	for _, link := range linksToVisit {
		wg.Add(1)
		go CheckLink(link, wg, t, &current, mut)
	}

	wg.Wait()
	fmt.Println("summary:")
	t.Render()
}
