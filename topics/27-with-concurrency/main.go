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

	/* Create channel */
	channel := make(chan string)

	/*
		Pass the channel to be used by the algorithm
	*/
	testServers(servers, channel)

	/*
		Print channel content (string) each time it is used (one time for every server)
	*/
	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}
}

/*
Receive the channel as parameter and pass it to the function with the logic
*/
func testServers(servers []string, channel chan string) {
	for _, server := range servers {
		/*
			Using the 'go' reserved word, all the calls will be executed simultaneously
		*/
		go httpChecker(server, channel)
	}
}

func httpChecker(url string, channel chan string) {
	_, err := http.Get(url)
	if err != nil {
		// Assign the desired result to the channel
		channel <- fmt.Sprintf("🔴 %s", url)
	} else {
		// Assign the desired result to the channel
		channel <- fmt.Sprintf("🟢 %s", url)
	}
}
