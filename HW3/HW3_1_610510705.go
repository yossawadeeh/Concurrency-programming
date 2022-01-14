package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// var m sync.Mutex
var buffer = make([]byte, 0, 10)

// var all sync.WaitGroup

var writter_ch = make(chan bool)
var consumer_ch = make(chan bool)

func main() {
	rand.Seed(10)
	// all.Add(1)
	go writer('a')
	go writer('b')
	go consumer()
	<-writter_ch
	<-consumer_ch
	// all.Wait()
}

func writer(c byte) {

	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Int63n(1e9)))
		// m.Lock()
		lb := len(buffer)
		if lb < cap(buffer) {
			buffer = buffer[:lb+1]
			buffer[lb] = c
			fmt.Printf("'%c' written to buffer.     buffer contents: %s\n",
				c, string(buffer))
		}
		// m.Unlock()
	}

	writter_ch <- true
}

func consumer() {
	a := []byte{'a'}
	b := []byte{'b'}

	for i := 0; i < 5; {
		time.Sleep(time.Duration(rand.Int63n(1e9)))
		// m.Lock()
		ai := bytes.Index(buffer, a)
		bi := bytes.Index(buffer, b)
		if ai >= 0 && bi >= 0 {
			if ai > bi {
				ai, bi = bi, ai
			}
			copy(buffer[bi:], buffer[bi+1:])
			copy(buffer[ai:], buffer[ai+1:])
			buffer = buffer[:len(buffer)-2]
			fmt.Printf("pair removed from buffer.  buffer contents: %s\n",
				string(buffer))
			i++
		}
		// m.Unlock()
	}

	consumer_ch <- true
	// all.Done()
}
