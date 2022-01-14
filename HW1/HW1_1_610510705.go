package main

import (
	"fmt"
	"sync"
)

func say(wg *sync.WaitGroup) {
	fmt.Println("goroutine: world")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		fmt.Println("main: hello")
		wg.Add(1)
		go say(&wg)
		wg.Wait()
	}

}
