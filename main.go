package main

import (
	"fmt"

	//"github.com/rishiagarwal-git/go-concurrency-patterns/intro/basic"
	"github.com/rishiagarwal-git/go-concurrency-patterns/intro/channel"
	//"github.com/rishiagarwal-git/go-concurrency-patterns/intro/goroutine"
)

func main(){
	// fmt.Println("Back to Basics!")
	// basic.Hello()
	// fmt.Println("Follow your 'Go-routine'!")
	// goroutine.Hello()
	fmt.Println("Channeling Go!")
	channel.Hello()
}

