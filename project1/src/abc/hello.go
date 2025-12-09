package main

import (
	"fmt"

)

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	ch2 <- "a"
	ch1 <- 1
	ch2 <- "b"
	ch1 <- 2

/*
	v1 := <-ch1
	v2 := <-ch2

	fmt.Println(v1)
	fmt.Println(v2)
	*/

	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	}
}