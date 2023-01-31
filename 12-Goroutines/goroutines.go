package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}
var counterMutex = 0

func main() {
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(100)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	var msg = "Hello"
	wg.Add(2)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHelloMutex()
		m.Lock()
		go incrementMutex()
	}
	go func(msg string) {
		fmt.Println(msg + "1")
		wg.Done()
	}(msg)
	go func() {
		fmt.Println(msg + "2")
		wg.Done()
	}()
	msg = "Goodbye"
	//time.Sleep(100 * time.Millisecond)
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

func sayHelloMutex() {
	fmt.Printf("HelloMutex #%v\n", counterMutex)
	m.RUnlock()
	wg.Done()
}

func incrementMutex() {
	counterMutex++
	m.Unlock()
	wg.Done()
}
