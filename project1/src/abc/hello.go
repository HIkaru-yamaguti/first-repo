package main

import (
	"fmt"

)

//スコープ


func appName() string {
	const AppName = "goaApp"
	var Version string = "1.0.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	//var b string = s
	b = s

	return b
}
func main() {
	fmt.Println("hello")

	fmt.Println(appName())

	fmt.Println(Do("Go言語"))
}