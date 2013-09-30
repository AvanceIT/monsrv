package main

import (
	"encoding/gob"
	"fmt"
	"github.com/AvanceIT/monitor/xmltools"
	"github.com/AvanceIT/monsrv/db"
	// "github.com/AvanceIT/monsrv/secure"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "192.168.0.5:2468")
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
			ev := xmltools.MonResult{}
			//secure.ClientAuth(c)
			dec := gob.NewDecoder(c)
			err = dec.Decode(&ev)
			if err != nil {
				fmt.Printf("decode error: %v\n", err)
			} else {
				db.AddEvent(ev)
			}
			c.Close()
		}(conn)
	}
}
