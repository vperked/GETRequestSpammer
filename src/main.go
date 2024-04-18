package main

import (
	"fmt"
	"net/http"
)

func main() {
	for i := 0; i < 5; i++ {
		get()
	}
}

func get() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	website := "https://vperked.online"
	resp, err := client.Get(website)
	resp.Header.Add("User-Agent", "Mozilla/5.0 (X11; CrOS x86_64 8172.45.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.64 Safari/537.36")
	if err != nil {
		fmt.Println("Could not send the GET Request = ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Spamming and it is responding with: ", resp.Status)
}
