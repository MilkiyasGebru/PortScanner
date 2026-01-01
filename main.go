package main

import (
	"PortScanner/port"
	"runtime"
)

func main() {

	portFinders := runtime.NumCPU()
	startPort := 10
	endPort := 1000
	done := make(chan interface{})
	portNumbersStream := port.GeneratePortNumber(done, startPort, endPort)
	finders := make([]<-chan int, portFinders)
	for i := 0; i < portFinders; i++ {
		finders[i] = port.ScanPortStream(done, portNumbersStream)
	}

}
