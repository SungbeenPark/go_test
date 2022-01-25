package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	people := [5]string{"been", "sung", "nico", "jerry", "abc"}
	for _, person := range people {
		go isSexy(person, c)
	}

	fmt.Println("Waiting for messages")
	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for ", i)
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
