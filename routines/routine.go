package main

import (
	"fmt"
	"time"
)

func main() {
	go backgroundrunner()
	go routineexample("hello")
	routineexample("world")

}
func routineexample(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(s)
	}
}
func backgroundrunner() {
	for {
		fmt.Println("runnning in the background")
		time.Sleep(2 * time.Second)
	}
}
