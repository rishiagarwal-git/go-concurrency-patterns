package select_

/* The fanin function in multiplexer has been re implemented using select for this example. */

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
	go func() { 
		for {
			select {
				case s := <-input1: c <-s
				case s := <-input2: c <-s
			}
		}
	}()
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