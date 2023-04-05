package main

import (
	"fmt"

	"github.com/rishiagarwal-git/go-concurrency-patterns/intro/basic"
	"github.com/rishiagarwal-git/go-concurrency-patterns/intro/channel"
	"github.com/rishiagarwal-git/go-concurrency-patterns/intro/goroutine"
	"github.com/rishiagarwal-git/go-concurrency-patterns/patterns/generator"
	"github.com/rishiagarwal-git/go-concurrency-patterns/patterns/multiplexer"
	"github.com/rishiagarwal-git/go-concurrency-patterns/patterns/restoring_sequence"
)

func main(){
	fmt.Println("Back to Basics!")
	basic.Hello()
	fmt.Println("Follow your 'Go-routine'!")
	goroutine.Hello()
	fmt.Println("Channeling Go!")
	channel.Hello()
	fmt.Println("Generating--->!!")
	generator.Hello()
	generator.Hello_v2()
	fmt.Println("Multiplexing!!")
	multiplexer.Hello()
	fmt.Println("REstoring Sequence pattern")
	restoring_sequence.Hello()
}

