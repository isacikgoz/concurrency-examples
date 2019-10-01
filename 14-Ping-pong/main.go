package main

import (
	"fmt"
	"time"
)

type ball struct {
	hits int
}

func main() {
	table := make(chan *ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(ball)
	time.Sleep(time.Second)
	<-table

}

func player(name string, table chan *ball) {
	for {
		ball := <-table // hits the table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
