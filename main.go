package main

import (
	"fmt"
	"net/http"
	"ioutil"
	"log"
)

var (
	i int
	timeout int = 10
)

func spam(url string) {
	for {
		http.Get(url)
	}
}
func spamget(url string) {
	go spam(url)
	fmt.Println("Get spamming started in the Backround")
}

func get(url string) {
	go ghettocount(timeout)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
   sb := string(body)
   log.Printf(sb)
   fmt.Println(sb)
}

func getnoreply(url string) {
	http.Get(url)
}

func getnoreplybackground(url string) {
	go getnoreply(url)
}

func ghettocount(num int) {
	i = 0
	for {
		count := i + 1
		i = count
		if i == num {
			goto done;
		}
	}
	done:
	fmt.Println("Timeout Exceeded Max Int")

}

func getfast(url string) {
	go get(url)
	fmt.Println("Request Sending In The Backround.")
}
