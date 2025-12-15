package main

import (
	"fmt"
	"os"

)

func main() {
	//Exit
	/*
	os.Exit(1)
	fmt.Println("start")
	*/
	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)
}
