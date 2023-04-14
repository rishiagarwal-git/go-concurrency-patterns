package timeout

import (
	"fmt"
	"math/rand"
	"time"
)

// timeout for each message
func Hello() {
	c := boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(800 * time.Millisecond):
			fmt.Println("You're too slow!")
			return
		}
	}
}

//timeout for complete Conversation
func Hello_v2() {
	c := boring("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You're too slow!")
			return
		}
	}
}

//timeout using Quit Channel
func Hello_v3() {
	quit := make(chan bool)
	c := boring_v3("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- { fmt.Println(<-c) }
	quit <-true
	fmt.Println("Stopped listening!")
}

func Hello_v4() {
	quit := make(chan string)
	c := boring_v4("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- { fmt.Println(<-c) }
	quit <-"bye"
	fmt.Printf("Joe Says: %q\n", <-quit)
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

func boring_v3(msg string, quit chan bool) <-chan string{ // Returns recieve-only channels of strings
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++{
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit: return
			}
		}
	}()
	return c
}

func boring_v4(msg string, quit chan string) <-chan string{ // Returns recieve-only channels of strings
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++{
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				cleanup()
				quit <-"See You"
				return
			}
		}
	}()
	return c
}

func cleanup() {
	fmt.Println("Cleaning up!")
}
