package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
	"github.com/playwright-community/playwright-go"
)

func IsValidURL(urlToCheck string) bool {
	parsed, err := url.Parse(urlToCheck)
	if err != nil {
		return false
	}

	if (parsed.Scheme != "http" && parsed.Scheme != "https") || parsed.Host == "" {
		return false
	}

	return true
}

func PopulateLinks(rootURL string, toPopulate *[]string) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v\n", err)
	}

	defer pw.Stop()

	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v\n", err)
	}

	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v\n", err)
	}

	defer page.Close()

	if _, err = page.Goto(rootURL); err != nil {
		log.Fatalf("could not goto: %s error: %v\n", rootURL, err)
	}

	pageLinks, err := page.Evaluate(`Array.from(document.querySelectorAll('a')).map(e => e.href)`)
	if err != nil {
		log.Fatalf("could not evaluate page: %v\n", err)
	}

	set := HashSet{}

	for _, link := range pageLinks.([]interface{}) {
		if IsValidURL(link.(string)) && !set.Has(link.(string)) {
			*toPopulate = append(*toPopulate, link.(string))
			set.Add(link.(string))
		}
	}
}

func colorStatus(response *http.Response) string {
	var status string

	if response.StatusCode >= 100 && response.StatusCode <= 199 {
		status = tml.Sprintf("<blue>%s<blue>", response.Status)

	} else if response.StatusCode >= 200 && response.StatusCode <= 299 {
		status = tml.Sprintf("<green>%s<green>", response.Status)

	} else if response.StatusCode >= 300 && response.StatusCode <= 399 {
		status = tml.Sprintf("<yellow>%s<yellow>", response.Status)

	} else {
		status = tml.Sprintf("<red>%s<red>", response.Status)
	}

	return status
}

func CheckLink(link string, wg *sync.WaitGroup, t *table.Table, current *int, mut *sync.RWMutex) {
	defer wg.Done()

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(link)

	if err != nil {
		fmt.Printf("an error occurred while visiting: %s error: %v\n", link, err)

		if urlErr, ok := err.(*url.Error); ok && urlErr.Timeout() {
			fmt.Printf("request to %s timed out\n", link)
		}

		return
	}

	defer resp.Body.Close()

	mut.Lock()
	t.AddRow(strconv.Itoa(*current), link, colorStatus(resp))
	*current += 1
	mut.Unlock()
}
