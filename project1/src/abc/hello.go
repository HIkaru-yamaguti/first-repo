package main

import (
	"fmt"

)
//カスタムエラー
//interface
/*
type error interface {
	Error() string
}*/

type Myerror struct {
	Message string
	ErrCode int
}

func (e *Myerror) Error() string {
	return e.Message
}

func RaiseError() error {
	return &Myerror{Message: "これはカスタムエラーです", ErrCode: 500}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	e, ok := err.(*Myerror)
	if ok {
		fmt.Println(e.ErrCode)
	}
}