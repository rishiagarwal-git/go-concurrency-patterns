package restoring_sequence

import (
	"fmt"
	"math/rand"
	"time"
)

//waitForIt := make(chan bool)

type Message struct {
	str string
	wait chan bool
}// shared between all messages

func Hello() {
	c := fanIn(boring("joe!"),boring("ann!"))// function returning a channel
	for i := 0; i < 5; i++ {
		msg1 := <-c; fmt.Println(msg1.str)
		msg2 := <-c; fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're both Boring I'm leaving")
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() { for {c <- <-input1}}()
	go func() { for {c <- <-input2}}()
	return c
}

func boring(msg string) <-chan Message{ // Returns recieve-only channels of strings
	c := make(chan Message)
	go func() { // We launch the goroutine from inside the function.
		waitForIt := make(chan bool) 
		for i := 0; ; i++ {
			c <- Message{ fmt.Sprintf("%s %d", msg, i), waitForIt} // Expression to be sent can be any value
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}