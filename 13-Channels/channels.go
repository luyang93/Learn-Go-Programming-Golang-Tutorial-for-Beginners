package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "Error"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

var logCh2 = make(chan logEntry, 50)
var doneCh2 = make(chan struct{})

func main() {
	ch := make(chan *int)
	wg.Add(2)
	go func() {
		i := <-ch
		time.Sleep(1 * time.Millisecond)
		fmt.Println(*i)
		wg.Done()
	}()
	go func() {
		i := 42
		ch <- &i
		i = 27
		wg.Done()
	}()

	ch2 := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch2
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch2 <- 42
			wg.Done()
		}()
	}

	ch3 := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch3
		fmt.Println(i)
		ch3 <- 273
		wg.Done()
	}()
	go func() {
		ch3 <- 423
		fmt.Println(<-ch3)
		wg.Done()
	}()

	ch4 := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch4)
	go func(ch chan<- int) {
		ch <- 424
		wg.Done()
	}(ch4)

	ch5 := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch5)
	go func(ch chan<- int) {
		ch <- 425
		ch <- 275
		close(ch)
		wg.Done()
	}(ch5)

	ch6 := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				fmt.Println(ok)
				break
			}
		}
		wg.Done()
	}(ch6)
	go func(ch chan<- int) {
		ch <- 426
		ch <- 276
		close(ch)
		wg.Done()
	}(ch6)

	go logger()
	defer func() {
		close(logCh)
	}()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}

	go logger2()
	logCh2 <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh2 <- logEntry{time.Now(), logInfo, "App is shutting down"}

	time.Sleep(100 * time.Millisecond)
	doneCh2 <- struct{}{}
	wg.Wait()
}

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}
}

func logger2() {
	for {
		select {
		case entry := <-logCh2:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh2:
			break
		default:
			break
		}

	}
}
