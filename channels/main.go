package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://github.com",
		"http://yahoo.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }
	//Equivalent to the same

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down.")
		c <- link
		return
	}
	fmt.Println(link, "is up and running")
	c <- link
}

/*
GO Routine is the current code in execution.
It is the current GO Routine.
To create A New GO ROUTINE which is the new thread for execution.
go funcCall(link)
*/
