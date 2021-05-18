package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://twitter.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// if only put go, the main routine could die off before the child routines could finish
		// hence, need to use channel for the main routine to "track" the child routines
		go checkLink(link, c)
	}

	// someone needs to wait for them to send something over
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-c, c)
	// 	i--
	// }

	// infinite loop
	// for {
	// 	go checkLink(<-c,c)
	// }

	for l := range c {

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // function literal example
	}

	// function literal --> lambda function
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

// to send data with channels

// channel <- X
// Y <- channel
// fmt.Println(<- channel)
