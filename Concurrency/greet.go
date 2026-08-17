package concurrency

import (
	"fmt"
	"time"
)

func GreetFunc(doneChan chan bool) {
	fmt.Println("Hello!!!!")
	doneChan <- true
}

func SlowGreet(doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello Guys!!!!")
	doneChan <- true
	close(doneChan)
}
