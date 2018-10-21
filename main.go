package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	for _, link := range links {
		checkLink(link)
	}
}

//this method will check the status of the links present in the slice one by one
//till the response is fetched, it won't be to the next link
func checkLink(link string) {
	// _is the response obj that comes along, second arg is the err
	_, err := http.Get(link)

	//if the error comes along
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up")
}
