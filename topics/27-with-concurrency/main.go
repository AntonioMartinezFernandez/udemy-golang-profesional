package main

import (
	"errors"
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
		"https://twitter.com",
		"https://www.wikipedia.org",
		"https://www.reddit.com",
		"https://www.anotherinventedurl.site",
		"https://www.linkedin.com",
	}

	/* Create channel */
	stringChannel := make(chan string)

	/*
		Pass the channel to be used by the algorithm
	*/
	testServers(servers, stringChannel)

	/*
		Print channel content (string) each time it is used (one time for every server)
	*/
	for range servers {
		fmt.Printf("%s\n", <-stringChannel)
	}
}

/*
Receive the channel as parameter and pass it to the function with the logic
*/
func testServers(servers []string, channel chan string) {
	for i := 0; i < len(servers); i++ {
		/*
			Using the 'go' reserved word, all the calls will be executed simultaneously
		*/
		go httpChecker(servers[i], channel)
	}
}

func httpChecker(url string, channel chan string) {
	tr := &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr, CheckRedirect: func(req *http.Request, via []*http.Request) error {
		// Allow a maximum of 15 redirects
		if len(via) >= 15 {
			return errors.New("too many redirects")
		}
		return nil
	}}

	req, reqError := http.NewRequest("GET", url, nil)
	if reqError != nil {
		panic(reqError)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://google.com/bot.html)")

	_, respError := client.Do(req)
	if respError != nil {
		// Assign the desired result to the channel
		channel <- fmt.Sprintf("🔴 %s - Error: %s", url, respError)
	} else {
		// Assign the desired result to the channel
		channel <- fmt.Sprintf("🟢 %s", url)
	}
}
