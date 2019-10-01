// Generator pattern: A function that returns a channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("boring!") // generating function

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring; I'm leaving")
}

func boring(msg string) <-chan string { // retruns receive-only channel of strings
	c := make(chan string)
	go func() { // launch go routine inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // return the channel
}
