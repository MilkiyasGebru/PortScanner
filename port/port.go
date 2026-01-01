package port

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func ScanPort(protocol string, hostname string, port int) bool {

	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 10*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true

}

func GeneratePortNumber(done <-chan interface{}, startPort, endPort int) <-chan int {

	portStream := make(chan int)

	go func() {
		defer close(portStream)
		for port := startPort; port <= endPort; port++ {
			select {
			case <-done:
				return
			case portStream <- port:
			}
		}
	}()

	return portStream

}

func ScanPortStream(done <-chan interface{}, portStream <-chan int) <-chan int {
	openPortStream := make(chan int)

	go func() {
		defer close(openPortStream)
		fmt.Println("Scanning Ports")
		for {
			//var port int

			select {
			case <-done:
				return
			case port, ok := <-portStream:
				if !ok {
					return
				}
				if ScanPort("tcp", "127.0.0.1", port) {
					openPortStream <- port
				}
			}
		}
	}()

	return openPortStream
}
