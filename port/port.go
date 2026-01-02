package port

import (
	"net"
	"strconv"
	"time"
)

func ScanPort(protocol string, hostname string, port int) bool {

	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 500*time.Millisecond)

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

func ScanPortStream(done <-chan interface{}, portStream <-chan int, hostname string) <-chan int {
	openPortStream := make(chan int)

	go func() {
		defer close(openPortStream)
		for {
			//var port int

			select {
			case <-done:
				return
			case port, ok := <-portStream:
				if !ok {
					return
				}
				if ScanPort("tcp", hostname, port) {
					openPortStream <- port
				}
			}
		}
	}()

	return openPortStream
}
