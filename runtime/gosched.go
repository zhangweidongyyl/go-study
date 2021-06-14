package main

import (
	"fmt"
	"runtime"
)

func say(str string) {
	fmt.Println(str)
}
func main() {
	go say("world")
	for i := 0; i < 3; i++ {
		say("hello")
		runtime.Gosched()
	}
}
