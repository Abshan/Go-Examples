package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	address := "example.com:80"
	timeout := 3 * time.Second

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Println("connect error:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("connected to %s\n", address)
}
