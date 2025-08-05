package main

import (
	"fmt"
	"time"
)

func Channels2() string {
	userch := make(chan string)

	for userch != nil {
		time.Sleep(2 * time.Second)
		person := <-userch
		fmt.Println(person)
	}
	userch <- "Arush"
	user := <-userch
	return user

}
