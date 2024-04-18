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
	if err != nil {
		fmt.Println("Could not send the GET Request = ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status", resp.Status)
}
