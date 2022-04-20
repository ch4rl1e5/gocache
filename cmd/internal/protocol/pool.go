package protocol

import "net"

var poolMap map[string]net.Conn

func pool(signature string, conn net.Conn) {
	poolMap[signature] = conn
}
