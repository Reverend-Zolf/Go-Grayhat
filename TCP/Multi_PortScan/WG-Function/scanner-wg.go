package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
	// Prevent the main function from exiting immediately
	var input string
	fmt.Scanln(&input)
}
