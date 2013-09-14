package main

import (
	"net"
	"fmt"
	"log"
	"github.com/AvanceIT/monsrv/secure"
)

func ExampleHello() {
	fmt.Println("Hello")
	// Output: Hello
}

func main() {
	l, err := net.Listen("tcp", ":2468")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			secure.ClientAuth(c)
			c.Close()
		}(conn)
	}
}
