package basic

import (
	"fmt"
	"math/rand"
	"time"
)

func Hello() {
	boring("Boring!")
}

func boring(msg string) {
	for i := 0; i <= 10; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	} 
}

func boringWithRandom(msg string) {
	for i := 0; i <= 10; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}