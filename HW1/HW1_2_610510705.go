package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 4; i++ {
		name("alice : ", "bob : ")
	}
}

func name(n1 string, n2 string) {
	n := make(chan string, 2)
	go func() { n <- n1 }()
	go func() { n <- n2 }()

	fmt.Printf(<-n)
	fmt.Println("ping")
	fmt.Printf(<-n)
	fmt.Println("pong")

}
