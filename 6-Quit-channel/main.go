// Quit channel: tell when to stop
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan bool)
	c := boring("boring!", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Printf("%s\n", <-c)
	}
	quit <- true
	fmt.Println("You are boring; I'm leaving")
}

func boring(msg string, quit chan bool) <-chan string { // retruns receive-only channel of strings
	c := make(chan string)
	go func() { // launch go routine inside the function
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // return the channel
}
