package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("concurrency\t")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCHI\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GORoutine\t", runtime.NumGoroutine())
	//wg.Add(1)
	//go foo()
	bar()
	//wg.Wait()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GORoutine\t", runtime.NumGoroutine())
	//overcome race condition use mutex

	var mu sync.Mutex
	//race condition
	count := 0
	const gs = 100
	var ws sync.WaitGroup
	ws.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := count //reading value do , writing same to share var
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			count = v
			mu.Unlock()
			ws.Done()
		}()
	}
	ws.Wait()

	fmt.Println("count", count)

	//atomic

	var counter int64

	ws.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Atomic counter\t", atomic.LoadInt64(&counter))
			runtime.Gosched()

			ws.Done()
		}()
	}
	ws.Wait()

	fmt.Println("counter", counter)

}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("Foo", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("Bar", i)
	}
}
