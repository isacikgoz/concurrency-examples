// Fan in pattern: A function that returns a channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("boring!"), boring("yay!")) // multiplexing function

	for i := 0; i < 9; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring; I'm leaving")
}

func fanIn(a, b <-chan string) <-chan string { // 2 go routines
	c := make(chan string)
	go func() {
		for {
			c <- <-a
		}
	}()
	go func() {
		for {
			c <- <-b
		}
	}()
	return c
}

func fanInSelect(a, b <-chan string) <-chan string { // single go routine
	c := make(chan string)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
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
