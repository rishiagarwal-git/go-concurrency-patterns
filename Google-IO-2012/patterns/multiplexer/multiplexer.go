package multiplexer

import (
	"fmt"
	"math/rand"
	"time"
)

func Hello() {
	c := fanIn(boring("joe!"),boring("ann!"))// function returning a channel
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both Boring I'm leaving")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for {c <- <-input1}}()
	go func() { for {c <- <-input2}}()
	return c
}

func boring(msg string) <-chan string{ // Returns recieve-only channels of strings
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any value
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}