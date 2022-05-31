package main

import (
	"bytes"
	"fmt"
	"github.com/ch4rl1e5/gocache/cmd/pkg/chunk"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	var chunks chunk.Chunk
	close := false
	for {
		buf := make([]byte, 8)
		for {
			size, err := c.Read(buf)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
				close = true
				break
			}

			if size == 0 {
				break
			}

			buf = bytes.TrimSuffix(buf[:size], []byte("\n"))
			if len(buf) == 0 {
				break
			}
			chunk.AppendChunk(buf, &chunks)
			break
		}
		if close {
			break
		}
	}
	c.Close()
}
