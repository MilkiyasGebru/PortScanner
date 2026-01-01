package main

import (
	"PortScanner/port"
	"fmt"
	"runtime"
	"sync"
)

func main() {

	portFinders := runtime.NumCPU()
	startPort := 10
	endPort := 60000
	done := make(chan interface{})
	portNumbersStream := port.GeneratePortNumber(done, startPort, endPort)
	finders := make([]<-chan int, portFinders)
	for i := 0; i < portFinders; i++ {
		finders[i] = port.ScanPortStream(done, portNumbersStream)
	}

	fanIn := func(done <-chan interface{}, channels ...<-chan int) <-chan int {
		var wg sync.WaitGroup
		wg.Add(len(channels))
		totalOpenPortStreams := make(chan int)
		combine := func(channel <-chan int) {
			defer wg.Done()
			for val := range channel {
				select {
				case <-done:
					return
				case totalOpenPortStreams <- val:
				}
			}
		}

		for _, channel := range channels {
			go combine(channel)
		}

		go func() {
			wg.Wait()
			close(totalOpenPortStreams)
		}()
		return totalOpenPortStreams
	}

	for openPort := range fanIn(done, finders...) {
		fmt.Println("Port Number ", openPort, "is open")
	}

}
