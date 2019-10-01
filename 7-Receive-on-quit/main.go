// Receive on quit channel: when do we know it's finished? Wait for it to tell us it's done.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan string)
	c := boring("boring!", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Printf("%s\n", <-c)
	}
	quit <- "bye!"
	fmt.Printf("You are boring; %q\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string { // returns receive-only channel of strings
	c := make(chan string)
	go func() { // launch go routine inside the function
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				cleanup()
				quit <- "see ya!"
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // return the channel
}

func cleanup() {
	fmt.Println("all clean.")
}
