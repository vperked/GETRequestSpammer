package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	for i := 0; i < 100; i++ {
		get()
	}
}

type Message struct {
	Content string `json:"content"`
}

func get() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	website := "https://inputwebsite"
	resp, err := client.Get(website)
	resp.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.3")
	if err != nil {
		fmt.Println("Could not send the GET Request = ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Sending GET requests too:", website, "and it is responding with: ", resp.Status)

	message := Message{
		Content: website + " " + resp.Status,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Couldnt marshal message! ", err)
		return
	}

	webhookURL := "https://discord.com/api/webhooks/1231044284464234537/X0U2CoW7jwhnaqrUCiifahbVV9jTm5CAILsyVdP394HXraE5RklbBfVwfmDPHqoaqskz"

	webhookResp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(messageJSON))
	if err != nil {
		fmt.Println("Couldnt send post request!", err)
		return
	}
	defer webhookResp.Body.Close()
	fmt.Println("Sent message")
}
