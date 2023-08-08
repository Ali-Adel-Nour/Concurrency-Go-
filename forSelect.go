package main

import "fmt"

//Pattern For Select

func main() {
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}
	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}
}
