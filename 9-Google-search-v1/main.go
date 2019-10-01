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

func Google(query string) (results []Result) {
	results = append(results, web(query))
	results = append(results, image(query))
	results = append(results, video(query))
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
