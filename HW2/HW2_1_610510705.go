// 1) นายภานุพันธ์ ตำปาน รหัสนักศึกษา 610510701	ตอนที่ 001
// 2) นายภูริวัฑฒก์ ชัยมงคล รหัสนักศึกษา 610510704 ตอนที่ 001
// 3) นางสาวยศวดี ห้อยมาลา รหัสนักศึกษา 610510705 ตอนที่ 001
// 4) นางสาวศศิวิมล วิทาทาน	รหัสนักศึกษา 610510707 ตอนที่ 001

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var sum = 0
var mu sync.Mutex

func counter(n int, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		mu.Lock()
		sum = sum + 1
		mu.Unlock()
	}
	fmt.Println("From ", n, ":", sum)
	wg.Done()
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter(i, &wg)
	}
	wg.Wait()
	fmt.Println("Final Sum:", sum)
}
