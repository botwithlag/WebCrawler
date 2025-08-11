package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	quit := make(chan int)
	wg.Add(2)
	go func() {
		for i := 0; i < 11; i++ {
			fmt.Print(<-c)
		}
		quit <- 0
		wg.Done()
	}()
	go fibo(quit, c)
	wg.Wait()
}
func routineexample(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(s)
	}
}
func backgroundrunner() {
	for {

		time.Sleep(2 * time.Second)
		fmt.Println("runnning in the background")
	}
}

func fibo(quit, c chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("lmao")
			wg.Done()
			return
		}
	}
}
