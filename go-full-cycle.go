package main

import (
	"fmt"
	"time"
)

type Course struct {
	Name        string
	Description string
	Price       int
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name :%s", c.Name)
}

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	channel := make(chan int)

	for i := 0; i < 10000; i++ {
		go worker(i, channel)
	}
}
