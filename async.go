package main

import (
	"fmt"
	"time"
)

func SomeFunc(num string) {
	fmt.Println(num)
}

func main() {
	go SomeFunc("2")
	go SomeFunc("3")
	go SomeFunc("4")
	time.Sleep(time.Second * 2)
	fmt.Println("hi")
}
