package main

import (
	"fmt"
)

func main() {
	var m = map[string]int{"A":100,"B":200}
	fmt.Println(m)

	m2 := map[string]int{"A":100,"B":200}
	fmt.Println(m2)

	m3 := map[int]string{
		1:"a",
		2:"b",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "japan"
	m4[2] = "usa"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[1])
	fmt.Println(m4[2], m3[1])

	s, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}
	fmt.Println(s, ok)

	m4[2] = "us"
	fmt.Println(m4)

	m4[3] = "china"
	fmt.Println(m4)

	fmt.Println(len(m4))

	delete(m4, 3)
	fmt.Println(m4)

	fmt.Println(len(m4))

	fmt.Println(m3)
}
