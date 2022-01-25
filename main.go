package main

import (
	"fmt"
	"time"
)

func main() {

	go sexCount("been")
	go sexCount("sung")
	time.Sleep(time.Second * 5)
}

func sexCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}

}
