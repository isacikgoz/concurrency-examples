// Timeout pattern: Wait timeout channel to return
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("boring!")

	to := time.After(3 * time.Second)
	for {
		select {
		case msg := <-c:
			fmt.Printf("%s\n", msg)
		case <-to:
			fmt.Println("You are too boring; I'm leaving")
			return
		}
	}
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
