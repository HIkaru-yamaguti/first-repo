package main

import (
	"fmt"
)

func main() {
	var ch1 chan int
/*
	var ch2 <-chan int

	var ch3 chan<- int

	fmt.Println(ch1, ch2, ch3)*/

	ch1 = make(chan int)

	ch2 := make(chan int)
	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <-1
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println(len(ch3))

	i := <- ch3
	fmt.Println(i)

	i2 := <- ch3
	fmt.Println(i2)
}
