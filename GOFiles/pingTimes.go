package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls  := [] string{"http://www.google.com", "http://www.yahoo.com", "http://www.bing.com"} //as an array
	start := time.Now()
	done  := make(chan string)

	for _, u := range urls {
		go func(u string) {
			resp, err := http.Get(u)
			if err != nil {
				done <- u + " " + err.Error()
			} else {
				done <- u + " " + resp.Status //var not func
			}
		}(u)
	}
	for _ = range urls {
		fmt.Println(<-done, time.Since(start))
	}
}
