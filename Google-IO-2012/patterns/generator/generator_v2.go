package generator

import (
	"fmt"
)

func Hello_v2() {
	joe := boring("joe!")// function returning a channel
	ann := boring("ann!")// function returning a channel
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both Boring I'm leaving")
}