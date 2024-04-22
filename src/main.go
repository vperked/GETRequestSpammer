package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Where the tom foolery begans.

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go get()
		defer wg.Done()
		get()
		wg.Wait()
		get()
	}
}

func get() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	website := "inputsite"
	resp, err := client.Get(website)
	resp.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.3")
	resp.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Safari/605.1.1")
	resp.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/125.0 Mobile/15E148 Safari/605.1.15")
	if err != nil {
		fmt.Println("Could not send the GET Request = ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Sending GET requests too:", website, "and it is responding with: ", resp.Status)
}
