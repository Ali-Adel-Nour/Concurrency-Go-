package main

import "fmt"

//Select: do lets a go routine wait on multiple commencations opeartions

func main() {

	myChannel := make(chan string)
	anotherChannel := make(chan string)

	//send data to the channel

	go func() {
		myChannel <- "data"
	}()
	go func() {
		anotherChannel <- "Ali"
	}()

	//Select statment will block until one of cases run
	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)

	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}

}
