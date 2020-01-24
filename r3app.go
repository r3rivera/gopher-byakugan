package main

import (
	"fmt"
	"gopher-byakugan/sms"
	"io/ioutil"
	"net/http"
	"runtime"
)

type urlInfo struct {
	url     string
	content int
}

func main() {

	fmt.Println("Start :: GO Routines -> ", runtime.NumGoroutine())

	sizeChannel := make(chan urlInfo)

	urls := []string{
		"https://example.com", "https://golang.org", "https://google.com",
	}

	for _, url := range urls {
		go responseSize(url, sizeChannel)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println("Channel Response ->", <-sizeChannel)
	}
	fmt.Println("End :: GO Routines -> ", runtime.NumGoroutine())
}

func responseSize(url string, sizeChannel chan urlInfo) {
	fmt.Println("Getting URL ->", url)

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("Error found! -> ", err)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error in body found! -> ", err)
		return
	}
	fmt.Println(sms.SendSMS())
	sizeChannel <- urlInfo{url: url, content: len(body)}
}
