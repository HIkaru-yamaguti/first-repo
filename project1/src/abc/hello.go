package main

import (
	"fmt"

)

type User struct {
	Name string
	Age  int
	//x, y int
}

func UpdateUser(user User) User {
	user.Name = "A"
	user.Age = 1000
	return user
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}
func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 30
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)
	
	user3 := User{Name: "user3", Age: 25}
	fmt.Println(user3)

	user4 := User{"user4", 28}
	fmt.Println(user4)

	//user5 := User{30, "user5"}
	//fmt.Println(user5)

	user6 := &User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7)

	user8 := &User{}
	fmt.Println(user8)

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)
}