package goroutine

import (
	"fmt"
	"math/rand"
	"time"
)

func Hello() {
	go boring("boring")
	fmt.Println("I'm Listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You're Boring I'm leaving")
}

func boring(msg string) {
	for i := 0; i <= 4; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}