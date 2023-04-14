package generator

import (
	"fmt"
	"math/rand"
	"time"
)

func Hello() {
	c := boring("boring!")// function returning a channel
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're Boring I'm leaving")
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