package main

import (
	"fmt"
	"github.com/AvanceIT/monsrv/db"
	"github.com/AvanceIT/monsrv/secure"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":2468")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	clients := db.ClientList()
	for _, line := range clients {
		fmt.Printf("%s\n", line)
	}

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
