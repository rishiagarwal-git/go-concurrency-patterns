package channel

import (
	"fmt"
	"math/rand"
	"time"
)

func Hello() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // recieve expression is just a value
	}
	fmt.Println("You're Boring I'm leaving")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any value
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}