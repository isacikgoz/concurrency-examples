// We want to get results from multiple backends
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	web   = fakeSearch("web")
	image = fakeSearch("image")
	video = fakeSearch("video")
)

type Result string

func main() {
	rand.Seed(time.Now().Unix())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)
}

// Google does not have any callbacks, locks, condition vars
func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- web(query) }()
	go func() { c <- image(query) }()
	go func() { c <- video(query) }()

	// with this we are waiting the slowest search result
	for i := 0; i < 3; i++ {
		results = append(results, <-c)
	}
	return
}

type Search func(query string) Result

// sleep for a while and return whatever the result is
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}
