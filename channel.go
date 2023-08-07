package main

import "fmt"

func main() {

	myChannel := make(chan string)

	//send data to the channel

	go func() {
		myChannel <- "data"
	}()

	//main function reading from the channel

	msg := <-myChannel

	fmt.Println(msg)
}
