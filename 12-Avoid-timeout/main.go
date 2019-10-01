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
	result := first("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"),
	)

	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}

// Google does not have any callbacks, locks, condition vars
func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- web(query) }()
	go func() { c <- image(query) }()
	go func() { c <- video(query) }()

	to := time.After(80 * time.Millisecond)
	// with this we are waiting the slowest search result
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-to:
			fmt.Println("timeout")
			return
		}
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

func first(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}
