// 1) นายภานุพันธ์ ตำปาน รหัสนักศึกษา 610510701	ตอนที่ 001
// 2) นายภูริวัฑฒก์ ชัยมงคล รหัสนักศึกษา 610510704 ตอนที่ 001
// 3) นางสาวยศวดี ห้อยมาลา รหัสนักศึกษา 610510705 ตอนที่ 001
// 4) นางสาวศศิวิมล วิทาทาน	รหัสนักศึกษา 610510707 ตอนที่ 001

package main

import (
	"fmt"
	"strconv"
)

var sum = 0

func counter(n int, c chan string) {

	for i := 0; i < 10000; i++ {
		sum = sum + 1
	}
	c <- "From " + strconv.Itoa(n) + " : " + strconv.Itoa(sum)

}

func main() {

	c := make(chan string)

	for i := 0; i < 5; i++ {
		go counter(i, c)
		fmt.Println(<-c)
	}

	fmt.Println("Final Sum:", sum)
}
