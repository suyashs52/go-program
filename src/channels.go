package main

import (
	"fmt"
	"sync"
)

//go run -race channels.go
func main() {
	fmt.Println("channels")

	// channel blocks , adding
	//c := make(chan int)

	// c <- 42

	// fmt.Println(<-c) // with same thread one is ready to put in
	// at same memory address another to take out at same time so deadlock
	//interlock so if another thread do take in and one thread do take out
	// it ll work

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	c = make(chan int, 1) //now we have buffer of 1 ,
	c <- 42
	//c <- 45 for buffer 1 we have 1 space only
	fmt.Println(<-c)
	fmt.Printf("\n%T\n", c)

	b := make(chan<- int, 2) //send them into only channel
	b <- 45

	//	d := make(<-chan int, 2)
	//	<-d //send from the channel

	//c=d specific didn't assign to general
	//d = c
	//general to specific converts
	fmt.Printf("%T\n", (<-chan int)(c))

	go sent(c)

	//go rece(c) it ll not block so nothing happen
	rece(c)

	fmt.Println("\nend")
	fmt.Println("even odd start")
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go sendData(even, odd, quit)

	receData(even, odd, quit)

	fmt.Println("even odd end")

	fanin := make(chan int)
	go sendDataCh(even, odd)
	go receiveCh(even, odd, fanin)

	for i := range fanin {
		fmt.Println("fanin", i)
	}

	ch := gen()

	for c := range ch {
		fmt.Println("receive from channel", c)
	}
}

func sendData(even chan int, odd chan int, quit chan<- int) {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}

	}

	//close(even)
	//close(odd)
	//quit <- 0
	close(quit)

}

func receData(even <-chan int, odd <-chan int, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("from the even channel", v)
		case v := <-odd:
			fmt.Println("from the odd channel", v)
		case i, v := <-quit:
			fmt.Println("from the quit channel", i, v)
			//	close(quit)
			return
		}
	}
}

//send

func sent(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)

}

//receive from channel
func rece(c <-chan int) {

	for i := range c {
		fmt.Println("receving ...", i)
	}

}

func sendDataCh(even chan<- int, odd chan<- int) {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}

	}

	close(even)
	close(odd)

}

func receiveCh(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := range even {
			fanin <- i
		}
		wg.Done()
	}()

	go func() {
		for i := range odd {
			fanin <- i
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)

}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 4; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
