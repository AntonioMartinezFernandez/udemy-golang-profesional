package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		fmt.Println("Total time: ", time.Since(startTime))
	}()

	servers := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.youtube.com",
		"https://www.antoniomartinezinventedurl.com",
		"https://www.twitter.com",
		"https://www.wikipedia.org",
		"https://www.reddit.com",
		"https://www.anotherinventedurl.site",
		"https://www.linkedin.com",
	}

	testServers(servers)
}

func testServers(servers []string) {
	for _, server := range servers {
		httpChecker(server)
	}
}

func httpChecker(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("🔴", url)
	} else {
		fmt.Println("🟢", url)
	}
}
