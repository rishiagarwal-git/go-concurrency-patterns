package main

import (
	"fmt"

	"github.com/rishiagarwal-git/go-concurrency-patterns/features/channel"
	"github.com/rishiagarwal-git/go-concurrency-patterns/features/goroutine"
	"github.com/rishiagarwal-git/go-concurrency-patterns/features/select_"
	"github.com/rishiagarwal-git/go-concurrency-patterns/intro/basic"
	daisy_chan "github.com/rishiagarwal-git/go-concurrency-patterns/patterns/daisy_chain"
	"github.com/rishiagarwal-git/go-concurrency-patterns/patterns/generator"
	"github.com/rishiagarwal-git/go-concurrency-patterns/patterns/multiplexer"
	"github.com/rishiagarwal-git/go-concurrency-patterns/patterns/restoring_sequence"
	"github.com/rishiagarwal-git/go-concurrency-patterns/patterns/timeout"
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
	fmt.Println("Restoring Sequence pattern")
	restoring_sequence.Hello()
	fmt.Println("Selecting----")
	select_.Hello()
	fmt.Println("Time is running out....")
	timeout.Hello()
	timeout.Hello_v2()
	timeout.Hello_v3()
	timeout.Hello_v4()
	fmt.Println("Daisy Chain Implemented")
	daisy_chan.Hello()
}

