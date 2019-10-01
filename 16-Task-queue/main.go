package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	numWorkers = 5
	numTasks   = 100
)

type task struct {
	id string
}

func main() {
	rand.Seed(time.Now().Unix())
	// create a buffered channel
	ch := make(chan task, 3)

	// run fixed number of workers
	for i := 0; i < numWorkers; i++ {
		go worker(ch)
	}

	hellaTasks := getTasks()
	processeds := 0

	for _, task := range hellaTasks {
		ch <- task
		processeds++
	}
	close(ch)
	fmt.Printf("processed total %d tasks\n", processeds)
}

func getTasks() []task {
	tasks := make([]task, 0)
	for i := 0; i < numTasks; i++ {
		tasks = append(tasks, task{id: strconv.Itoa(rand.Intn(1e9))})
	}
	return tasks
}

func worker(ch chan task) {
	for {
		t := <-ch // receive task
		time.Sleep(time.Duration(rand.Intn(5e2)) * time.Millisecond)
		process(t)
	}
}

func process(t task) {
	fmt.Printf("processing TASK: %s\n", t.id)
}
