package main

import (
	"io"
	"log"
	"net"
	"runtime"
	"strings"
)

func tcpServe(listener net.Listener) {
	log.Printf("TCP: listening on %s", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				log.Printf("NOTICE: temporary Accept() failure - %s", err.Error())
				runtime.Gosched()
				continue
			}

			// theres no direct way to detect this error because it is not exposed
			if !strings.Contains(err.Error(), "use of closed network connection") {
				log.Printf("ERROR: listener.Accept() - %s", err.Error())
			}
			break
		}

		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}

	log.Printf("TCP: closing %s", listener.Addr().String())
}
