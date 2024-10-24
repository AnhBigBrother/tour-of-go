package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

type SafeCrawl struct {
	mux     sync.Mutex
	crawled map[string]bool
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.

func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	// Solution1: use WaitGroup to force each goroutines wait for all child-goroutine before exiting.

	var wg sync.WaitGroup
	sc := SafeCrawl{crawled: map[string]bool{}}
	var CrawlRecursive func(url string, depth int, wg *sync.WaitGroup)

	CrawlRecursive = func(url string, depth int, wg *sync.WaitGroup) {
		sc.mux.Lock()
		if depth <= 0 || sc.crawled[url] {
			sc.mux.Unlock()
			wg.Done()
			return
		}
		sc.crawled[url] = true
		sc.mux.Unlock()
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		var wg_child sync.WaitGroup
		for _, u := range urls {
			wg_child.Add(1)
			go CrawlRecursive(u, depth-1, &wg_child)
		}
		wg_child.Wait()
		wg.Done()
	}
	wg.Add(1)
	go CrawlRecursive(url, depth, &wg)
	wg.Wait()
}

func Crawl2(url string, depth int, fetcher Fetcher) {
	// Solution2: use Chanel to force each goroutines wait for all child-goroutine before exiting.
	exit := make(chan byte)
	sc := SafeCrawl{crawled: map[string]bool{}}
	var CrawlRecursive func(url string, depth int, exit chan byte)

	CrawlRecursive = func(url string, depth int, exit chan byte) {
		sc.mux.Lock()
		if depth <= 0 || sc.crawled[url] {
			sc.mux.Unlock()
			exit <- 1
			return
		}
		sc.crawled[url] = true
		sc.mux.Unlock()
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			exit <- 1
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		exit_child := make(chan byte)
		for _, u := range urls {
			go CrawlRecursive(u, depth-1, exit_child)
		}
		for i := 0; i < len(urls); i++ {
			<-exit_child
		}
		close(exit_child)
		exit <- 1
	}
	go CrawlRecursive(url, depth, exit)
	<-exit
	close(exit)
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}
