// Restoring sequencing: send a channel on a channel making goroutine wait its turn
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type message struct {
	str string

	// act as a signaler, it will block on the wait channel until
	// other person says OK I want you to go ahead
	wait chan bool
}

func main() {
	c := fanIn(boring("boring!"), boring("yay!")) // multiplexing function

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You are boring; I'm leaving")
}

func fanIn(a, b <-chan message) <-chan message {
	c := make(chan message)
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

func boring(msg string) <-chan message { // retruns receive-only channel of strings
	c := make(chan message)
	waitForIt := make(chan bool) // shared between all messages
	go func() {                  // launch go routine inside the function
		for i := 0; ; i++ {
			c <- message{str: fmt.Sprintf("%s %d", msg, i), wait: waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c // return the channel
}
