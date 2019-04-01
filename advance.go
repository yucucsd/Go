package main

import (
	"fmt"
	"time"
)

type phone interface {
	call()
}

type nokiaphone struct {

}

func (nokiaPhone nokiaphone) call() {
	fmt.Println("Nokia")
}

type iphone struct {

}

func (Iphone iphone) call() {
	fmt.Println("Iphone")
}

func say(s string) {
	for i:= 0; i < 5; i++ {
		time.Sleep(10 *time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v 
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c) /* deadlock if we do not close the channel */
}

func main() {
	var ph phone
	ph = new(nokiaphone)
	ph.call()
	ph = new(iphone)
	ph.call()
	go say("hello")
	say("world")
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //receive from channel c
	fmt.Println(x, y, x + y)

	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1)
	for i := range c1 {
		fmt.Println(i)
	}
}
