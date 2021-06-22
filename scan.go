package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var name string
	fmt.Print("Enter the address: ")
	fmt.Scanf("%s\n", &name)

	for i := 1; i <= 1024; i++ {
		time.Sleep(30 * time.Millisecond)
		go func(j int) {
			address := fmt.Sprintf("%s:%d", name, j)

			timeout := net.Dialer{Timeout: 1 * time.Second}
			conn, err := timeout.Dial("tcp", address)
			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
}
