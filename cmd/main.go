package cmd

import (
	"github.com/ch4rl1e5/gocache/cmd/internal/protocol"
)

func main() {

	server := protocol.NewServer()
	server.Listen(":8081")
	defer server.Close()
}
