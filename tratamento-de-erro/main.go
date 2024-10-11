package main

import (
	"fmt"
)

type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error: %s, code: %d", e.Message, e.Code)
}

func ErrorTest(msg string, code int) error {
	return &MyError{Message: msg, Code: code}
}

func main() {
	err := ErrorTest("testando erro", -1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Jamais irei rodar :/")
}
