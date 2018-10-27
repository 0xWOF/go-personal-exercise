package main

import (
	"time"
)

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	close(ch)

	println(<-ch)
	// println(<-ch)

	if _, success := <-ch; !success {
		println(success)
		println("no more data")
	}

	//-

	done1 := make(chan bool)
	done2 := make(chan bool)

	go run2(done2)
	go run1(done1)

	time.Sleep(1 * time.Second)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")
			break EXIT

		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}

	go run2(done2)
	go run1(done1)

	time.Sleep(1 * time.Second)

EXIT2:
	for {
		select {
		case <-done2:
			println("run2 완료")
			break EXIT2
		case <-done1:
			println("run1 완료")
			break EXIT2
		}
	}
}

func run1(done chan bool) {
	done <- true
}

func run2(done chan bool) {
	done <- true
}
