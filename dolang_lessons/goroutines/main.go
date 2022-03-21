package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sayHello(ch)

	for i := range ch {

		fmt.Println(i)
	}

}

func Say(word string) {
	//time.Sleep(1 * time.Second)
	fmt.Println(word)

}

func sayHello(exit chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		Say("Hello")
		exit <- "exit"
	}
	close(exit)
}
