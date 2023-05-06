package main

import "fmt"

func main() {
	channel1 := make(chan string, 1)
	go func(ch <-chan string) {
		val, bool1 := <-channel1
		fmt.Println(val)
		fmt.Println(bool1)

	}(channel1)

	go func(ch <-chan string) {
		close(channel1)
		channel1 <- "Hello Dhebug"
	}(channel1)

}
