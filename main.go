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

	//creating a new channel
	c := make(chan string)

	/*
			when we use "go" infront a function call, it means run the function inside runs on a brand new go routine
		And when that happens, and the code running inside the new Go Routine encounters a blockage, the control of the program goes back to the last Go routine.


	*/

	//necessary to pass the channel to the checklink function as it is available only inside the main function and the checklink function would be able to use it only when the channel is passed to it

	for _, link := range links {
		go checkLink(link, c)
	}

	//receiving the message sent from the channel

	//creating an infinte for loop
	for {
		//receiving is a blocking operation
		//first arg for channel, second arg for type
		go checkLink(<-c, c)
	}

}

//this method will check the status of the links present in the slice one by one
//till the response is fetched, it won't be to the next link
func checkLink(link string, c chan string) {
	// _is the response obj that comes along, second arg is the err
	_, err := http.Get(link)

	//if the error comes along
	if err != nil {
		fmt.Println(link, "might be down!")
		//pushing the link to the channel
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link
}
